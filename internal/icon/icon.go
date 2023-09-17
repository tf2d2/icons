/*
Copyright © 2023 The tf2d2 Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package icon

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/hashicorp/go-hclog"
)

var (
	sources = []string{
		"https://aws.amazon.com/architecture/icons/",
		"https://cloud.google.com/icons",
	}

	downloadUrls = map[string]string{
		// the source site provides the download url after clicking a html checkbox
		// which cannot be done programmatically with colly, so the download url
		// is manually retrieved and copied here
		// this manual task should be deprecated by using browser automation, e.g. selenium
		// ref: https://learn.microsoft.com/en-us/azure/architecture/icons/
		"azure": "https://arch-center.azureedge.net/icons/Azure_Public_Service_Icons_V17.zip",
	}

	awsLinkRgx   = regexp.MustCompile(`/Asset-Package_([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)\.zip$`)
	gcpLinkRgx   = regexp.MustCompile(`google-cloud-icons\.zip$`)
	removePrefix = regexp.MustCompile(`(Azure_Public_Service_Icons/Icons)|([\d]+-icon-service-)`)
	removeSuffix = regexp.MustCompile(`(_[\d]{8})`)

	awsReplacer = strings.NewReplacer(
		"Architecture-Service-Icons", "service", "Arch_", "",
		"Category-Icons", "category", "Arch-Category_", "",
		"Resource-Icons", "resource", "Res_", "",
		":", "-", // to prevent malformed file paths when importing as module
		"_16", "", "_32", "", "_48", "", "_64", "", "48_", "", // order matters, this must be last
	)
)

// GetIconDetails finds and downloads the architecture icons for all relevant cloud providers
func GetIconDetails(ctx context.Context) error {
	logger := hclog.FromContext(ctx)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		logger.Info(fmt.Sprintf("searching for download url at: %s", r.URL))
	})

	var collyErr error
	c.OnError(func(_ *colly.Response, err error) {
		collyErr = err
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if awsLinkRgx.MatchString(link) {
			downloadUrls["aws"] = link
			logger.Info(fmt.Sprintf("found download url: %s", link))
		} else if gcpLinkRgx.MatchString(link) {
			link = fmt.Sprintf("https://%s", e.Attr("track-metadata-eventdetail"))
			downloadUrls["gcp"] = link
			logger.Info(fmt.Sprintf("found download url: %s", link))
		}
	})

	for _, url := range sources {
		err := c.Visit(url)
		if err != nil {
			return fmt.Errorf("error visiting: %s", url)
		}
	}

	for cloud, url := range downloadUrls {
		outputFile := fmt.Sprintf("%s.zip", cloud)

		err := downloadFile(ctx, url, outputFile)
		if err != nil {
			return err
		}

		err = unzipFile(ctx, outputFile)
		if err != nil {
			return err
		}

		err = removeZipfile(ctx, outputFile)
		if err != nil {
			return err
		}
	}

	return collyErr
}

func downloadFile(ctx context.Context, link, file string) error {
	logger := hclog.FromContext(ctx)

	// open the output file for writing
	output, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer output.Close() //nolint:errcheck

	// download zip file
	parsedLink, err := url.Parse(link)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedLink.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating download request: %s", err.Error())
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading file %s: %s", file, err.Error())
	}
	defer response.Body.Close() //nolint:errcheck

	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("HTTP error: %s", response.Status))
	}

	// Copy the response body to the output file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("error creating output file: %s", err.Error())
	}

	logger.Info(fmt.Sprintf("downloaded %s to %s", filepath.Base(link), file))

	return nil
}

func unzipFile(ctx context.Context, file string) error {
	logger := hclog.FromContext(ctx)

	dst := strings.TrimSuffix(file, filepath.Ext(file))

	// open the zip file
	archive, err := zip.OpenReader(file)
	if err != nil {
		return fmt.Errorf("error opening zip file: %s", err.Error())
	}
	defer archive.Close() //nolint:errcheck

	// extract the files from the zip
	for _, f := range archive.File {
		// skip files
		if skipFiles(f.Name) {
			continue
		}

		// create the destination file path
		filePath := filepath.Clean(getFilepath(dst, f.Name))

		// check if the file is a directory
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return fmt.Errorf("error creating directory: %s", err.Error())
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("error creating parent directory: %s", err.Error())
		}

		// Create an empty destination file
		dstFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating empty destination file: %s", err.Error())
		}

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("error opening source file: %s", err.Error())
		}
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return fmt.Errorf("error copying to destination file: %s", err.Error())
		}

		// Close the files
		_ = dstFile.Close()
		_ = srcFile.Close()
	}

	logger.Info(fmt.Sprintf("unzipped %s icons", file))

	return nil
}

func skipFiles(name string) bool {
	// skip __MACOSX directory
	if strings.HasPrefix(name, "__MACOSX") {
		return true
	}

	// skip .DS_Store files
	if strings.Contains(name, ".DS_Store") {
		return true
	}

	// skip PNG files
	if filepath.Ext(name) == ".png" {
		return true
	}

	// skip Azure PDF files
	if filepath.Ext(name) == ".pdf" {
		return true
	}

	// skip specific AWS icon sizes
	// if strings.Contains(name, "16") {
	// 	return true
	// }
	// if strings.Contains(name, "32") {
	// 	return true
	// }
	// if strings.Contains(name, "48") {
	// 	return true
	// }

	return false
}

func getFilepath(dst, file string) string {
	// remove prefixes
	name := removePrefix.ReplaceAllString(file, "")

	// remove suffixes
	name = removeSuffix.ReplaceAllString(name, "")

	// replace original directory names
	name = awsReplacer.Replace(name)

	// create the destination file path
	filePath := filepath.Join(dst, name)

	return filePath
}

func removeZipfile(ctx context.Context, filename string) error {
	logger := hclog.FromContext(ctx)

	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info(fmt.Sprintf("file %s does not exist", filename))
		} else {
			return err
		}
	} else {
		err = os.RemoveAll(filename)
		if err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("removed file: %s", filename))
	}

	return nil
}
