package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tf2d2/icons/internal/icon"
	"github.com/tf2d2/icons/internal/resource"
	"github.com/tf2d2/icons/internal/template"
	"github.com/tf2d2/icons/internal/terraform"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	outputDir = "providers"
)

var (
	providers = []string{"aws", "azurerm", "google"}
)

// Runtime represents the CLI execution runtime
type Runtime struct {
	ctx    context.Context
	cmd    *cobra.Command
	Config *config
}

// config represents the CLI runtime config
type config struct {
	CfgFile string `mapstructure:"-"`
	Verbose bool   `mapstructure:"verbose"`
	DryRun  bool   `mapstructure:"dry-run"`
}

// NewCLIRuntime returns a new CLI runtime instance
func NewCLIRuntime(ctx context.Context) *Runtime {
	return &Runtime{
		cmd:    nil,
		ctx:    ctx,
		Config: &config{},
	}
}

// PreRunE executes before `RunE` to configure logging, etc.
func (r *Runtime) PreRunE(cmd *cobra.Command, _ []string) error {
	logger := hclog.FromContext(r.ctx)

	r.cmd = cmd

	if err := r.bindFlags(); err != nil {
		return err
	}

	if err := r.readConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(r.Config); err != nil {
		return err
	}

	if r.Config.Verbose {
		logger.SetLevel(hclog.LevelFromString("debug"))
	}

	return nil
}

// RunE executes the logic to generate architecture icon details
func (r *Runtime) RunE(_ *cobra.Command, _ []string) error {
	return r.generateContent()
}

// readConfig loads Viper config
func (r *Runtime) readConfig() error {
	logger := hclog.FromContext(r.ctx)

	if r.cmd.Flags().Changed("config") {
		viper.SetConfigFile(r.Config.CfgFile)
	} else {
		viper.SetConfigName(".tf2d2-icons")
		viper.SetConfigType("yml")
	}

	// Search config in the following directories with name ".tf2d2-icons" (without extension)
	viper.AddConfigPath("$HOME/") // home directory
	viper.AddConfigPath(".")      // current directory

	// If a config file is found, read it in.
	err := viper.ReadInConfig()

	var notFoundError viper.ConfigFileNotFoundError
	switch {
	case err != nil && errors.As(err, &notFoundError):
		logger.Warn("no config file found", "warning", err.Error())
	case err != nil && !errors.As(err, &notFoundError):
		// pathError and notFoundError are produced when no config file is found
		// here we check and return an error produced when reading the config file
		logger.Error("failed to read config", "error", err.Error())
		return err
	default:
		logger.Info("using config file", "path", viper.ConfigFileUsed())
	}
	return nil
}

// bindFlags binds all Cobra flags to Viper config
func (r *Runtime) bindFlags() error {
	fs := r.cmd.Flags()
	return viper.BindPFlags(fs)
}

func (r *Runtime) generateContent() error {
	logger := hclog.FromContext(r.ctx)

	// get icon details
	err := icon.GetIconDetails(r.ctx)
	if err != nil {
		logger.Error("error retrieving icon details", "error", err.Error())
	}

	// get provider resource details
	resources, err := terraform.GetProviderResourceDetails(r.ctx)
	if err != nil {
		return err
	}

	// check output directories
	err = checkOutputDirectories(r.ctx)
	if err != nil {
		return err
	}

	// generate json files
	err = generateJSONFiles(r.ctx, resources)
	if err != nil {
		return err
	}

	// render provider resource templates
	err = writeTemplates(r.ctx, resources)
	if err != nil {
		return err
	}

	return nil
}

func checkOutputDirectories(ctx context.Context) error {
	logger := hclog.FromContext(ctx)

	for _, c := range providers {
		path := filepath.Join(outputDir, c)
		_, err := os.Stat(path)
		switch {
		case os.IsNotExist(err):
			err = os.MkdirAll(path, 0750)
			if err != nil {
				return fmt.Errorf("error creating output directory: %s", err.Error())
			}
			logger.Info(fmt.Sprintf("created output directory: %s", path))
		case err != nil:
			return fmt.Errorf("error checking output directory: %s", err.Error())
		}
	}

	return nil
}

//nolint:gocyclo // simplify this function
func generateJSONFiles(ctx context.Context, resources map[string][]*resource.Resource) error {
	logger := hclog.FromContext(ctx)

	for k, v := range resources {
		jsonFilename := filepath.Clean(filepath.Join(outputDir, k, strings.ToLower(k)+".json"))
		data, err := os.ReadFile(jsonFilename)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("error reading json file %s: %s", jsonFilename, err.Error())
		}

		currentJSONResources := make(map[string]string)
		if len(data) > 0 {
			err = json.Unmarshal(data, &currentJSONResources)
			if err != nil {
				return fmt.Errorf("error unmarshalling %s: %s", jsonFilename, err.Error())
			}
		}

		current := len(currentJSONResources)
		added := 0
		removed := 0

		for _, r := range v {
			if _, exists := currentJSONResources[r.Name]; !exists {
				currentJSONResources[r.Name] = ""
				added++
			}
		}

		// delete resource no longer found in provider
		if current > 0 {
			for ck, cv := range currentJSONResources {
				found := false
				for _, r := range v {
					if ck == r.Name {
						found = true
						r.IconURL = cv
					}
				}
				if !found {
					removed++
					delete(currentJSONResources, ck)
				}
			}

			// delete existing json file
			err = os.Remove(jsonFilename)
			if err != nil {
				return fmt.Errorf("error deleting file %s: %s", jsonFilename, err.Error())
			}
		}

		// re-create json file with updated content
		f, err := os.OpenFile(jsonFilename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("error opening file %s: %s", jsonFilename, err.Error())
		}
		defer f.Close()

		e := json.NewEncoder(f)
		e.SetEscapeHTML(false)
		e.SetIndent("", "  ")
		err = e.Encode(currentJSONResources)
		if err != nil {
			return fmt.Errorf("error encoding %s icon data: %s", jsonFilename, err.Error())
		}

		logger.Info(fmt.Sprintf("generated %s", jsonFilename),
			"before", current, "after", current+added-removed, "added", added, "removed", removed)
	}

	return nil
}

func writeTemplates(ctx context.Context, resources map[string][]*resource.Resource) error {
	logger := hclog.FromContext(ctx)

	for k, v := range resources {
		td := template.Template{
			Provider:  strings.ToUpper(k),
			Resources: v,
		}
		name := strings.ToLower(k)
		filename := fmt.Sprintf("%s.go", name)
		filepath := filepath.Join(outputDir, name, filename)
		err := template.Render(filepath, td)
		if err != nil {
			return fmt.Errorf("error writing template: %s", err.Error())
		}
		logger.Info(fmt.Sprintf("generated %s", filepath))
	}

	return nil
}
