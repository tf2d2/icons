package aws

import (
	"fmt"

	"github.com/tf2d2/icons/internal/resource"
)

var (
	Resources = []*resource.Resource{
  {{- range .Resources }}
		{
			Name:        "{{ .Name }}",
			{{- if ne .IconURL "" }}
			IconURL: 	 "{{ .IconURL }}",
			{{- end }}
		},
  {{- end }}
	}
	resourcesMap = map[string]int{
		{{- range $index, $el := .Resources }}
					"{{ .Name }}": {{ $index }},
		{{- end }}
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
