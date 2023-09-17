package terraform

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tf2d2/icons/internal/resource"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-exec/tfexec"
	"golang.org/x/exp/slices"
)

const (
	workDir                = "internal/terraform"
	providerRegistryPrefix = "registry.terraform.io/hashicorp/"
)

// GetProviderResourceDetails returns the resource details for all relevant providers
func GetProviderResourceDetails(ctx context.Context) (map[string][]*resource.Resource, error) {
	logger := hclog.FromContext(ctx)

	path, err := exec.LookPath("terraform")
	if err != nil {
		return nil, fmt.Errorf("error finding terraform binary, %s", err.Error())
	}

	tf, err := tfexec.NewTerraform(filepath.Clean(workDir), path)
	if err != nil {
		return nil, fmt.Errorf("error running NewTerraform: %s", err.Error())
	}

	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		return nil, fmt.Errorf("error running `terraform init`: %s", err.Error())
	}

	providerSchemas, err := tf.ProvidersSchema(ctx)
	if err != nil {
		return nil, fmt.Errorf("error running `terraform providers schema -json`: %s", err.Error())
	}

	resources := make(map[string][]*resource.Resource)
	for pk, pv := range providerSchemas.Schemas {
		pName := strings.ReplaceAll(pk, providerRegistryPrefix, "")
		pURL := pk
		pResourceCount := len(pv.ResourceSchemas)
		logger.Info(fmt.Sprintf("provider: %s, address: %s, total resources: %s", pName, pURL, strconv.Itoa(pResourceCount)))

		var resourceNames []string
		for rk := range pv.ResourceSchemas {
			resourceNames = append(resourceNames, rk)
		}
		slices.Sort(resourceNames)

		for _, i := range resourceNames {
			r := &resource.Resource{
				Name: i,
			}
			resources[pName] = append(resources[pName], r)
		}
	}

	logger.Info("generated terraform provider schema details")

	return resources, nil
}
