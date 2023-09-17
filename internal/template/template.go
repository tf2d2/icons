package template

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tf2d2/icons/internal/resource"
)

//go:embed resources.tmpl
var tpl string

type Template struct {
	Provider  string
	Resources []*resource.Resource
}

func New(provider string, resources []*resource.Resource) *Template {
	return &Template{
		Provider:  provider,
		Resources: resources,
	}
}

// Render renders provider resource details to output file
func Render(filename string, td Template) error {
	f, err := os.OpenFile(filepath.Clean(filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file (%s): %s", filename, err.Error())
	}
	defer f.Close() //nolint:errcheck

	tplate, err := template.New("resource").Parse(tpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %s", err.Error())
	}

	var buffer bytes.Buffer
	err = tplate.Execute(&buffer, td)
	if err != nil {
		return fmt.Errorf("error executing template: %s", err.Error())
	}

	b, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting generated source file")
	}
	_, _ = io.Copy(f, bytes.NewBuffer(b))

	return nil
}
