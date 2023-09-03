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

package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

var (
	sources = []string{
		"https://aws.amazon.com/architecture/icons/",
		"https://cloud.google.com/icons",
	}

	downloadUrls = map[string]string{
		"azure": "https://arch-center.azureedge.net/icons/Azure_Public_Service_Icons_V17.zip",
	}

	awsLinkRgx   = regexp.MustCompile(`/Asset-Package_([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)\.zip$`)
	gcpLinkRgx   = regexp.MustCompile(`google-cloud-icons\.zip$`)
	removeSuffix = regexp.MustCompile(`(_[\d]{8})`)
	removePrefix = regexp.MustCompile(`(Azure_Public_Service_Icons/Icons)`)

	awsReplacer = strings.NewReplacer(
		"Architecture-Service-Icons", "service", "Arch_", "",
		"Category-Icons", "category", "Arch-Category_", "",
		"Resource-Icons", "resource", "Res_", "",
		"_16", "", "_32", "", "_48", "", "_64", "", // order matters, this must be last
	)
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("scraping: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("error:", err)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if awsLinkRgx.MatchString(link) {
			downloadUrls["aws"] = link
		} else if gcpLinkRgx.MatchString(link) {
			downloadUrls["gcp"] = fmt.Sprintf("https://%s", e.Attr("track-metadata-eventdetail"))
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("completed:", r.Request.URL)
	})

	for _, url := range sources {
		err := c.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	for cloud, url := range downloadUrls {
		log.Printf("downloading %s icons: %s", cloud, url)
		outputFile := fmt.Sprintf("%s.zip", cloud)
		downloadFile(url, outputFile)
		unzipFile(outputFile)
	}
}

func downloadFile(link, file string) {
	// delete existing zip output file
	checkFileExists(file)

	// open the output file for writing
	output, err := os.Create(filepath.Clean(file))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = output.Close()
		if err != nil {
			panic(err)
		}
	}()

	// download zip file
	ctx := context.Background()
	parsedLink, err := url.Parse(link)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedLink.String(), nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("HTTP error: %s", response.Status))
	}

	// Copy the response body to the output file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		panic(err)
	}

	log.Println("saved", filepath.Base(link), "to", file)
}

func unzipFile(file string) {
	dst := strings.TrimSuffix(file, filepath.Ext(file))
	log.Printf("unziping %s icons", dst)

	// delete existing unzip output folder
	checkFileExists(dst)

	// open the zip file
	log.Println("reading", file)
	archive, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalf("error opening zip file: %s", err)
	}
	defer closeZipFile(archive)

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
				panic(fmt.Errorf("error creating directory: %w", err))
			}
			continue
		}

		// Create the parent directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(fmt.Errorf("error creating parent directory: %w", err))
		}

		// Create an empty destination file
		dstFile, err := os.Create(filePath)
		if err != nil {
			panic(fmt.Errorf("error creating empty destination file: %w", err))
		}

		// Open the file in the zip and copy its contents to the destination file
		srcFile, err := f.Open()
		if err != nil {
			panic(fmt.Errorf("error opening source file: %w", err))
		}
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			panic(fmt.Errorf("error copying to destination file: %w", err))
		}

		// Close the files
		_ = dstFile.Close()
		_ = srcFile.Close()
	}
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

	// skip AWS icon sizes < 64
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

func checkFileExists(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file %s does not exist", filename)
		} else {
			log.Fatal(err)
		}
	} else {
		err = os.RemoveAll(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("removed file: %s", filename)
	}
}

func closeZipFile(f *zip.ReadCloser) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
