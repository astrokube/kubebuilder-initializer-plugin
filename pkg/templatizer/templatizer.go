package templatizer

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/astrokube/layout-kubebuilder-plugin/pkg/templatizer/internal/source"
)

func Templatize(sourceType string, connString string, vars interface{}) (map[string]string, error) {
	var content map[string]string
	var err error
	if sourceType == "" || strings.EqualFold(sourceType, "git") {
		content, err = source.NewGitSource(connString).GetTemplateContent()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsupported source '%s'", sourceType)
	}
	for k, v := range content {
		processedFile, err := processTemplateFile(k, v, vars)
		if err != nil {
			return nil, err
		}
		content[k] = processedFile
	}
	return content, nil
}

func processTemplateFile(name string, content string, variables interface{}) (string, error) {
	t, err := template.New(name).Parse(content)
	if err != nil {
		return "", fmt.Errorf("error parsing file: '%w'", err)
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, variables); err != nil {
		return "", fmt.Errorf("error processing template: %w", err)
	}
	return buf.String(), nil
}
