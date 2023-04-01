package templatizer

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/astrokube/layout-kubebuilder-plugin/pkg/templatizer/internal/source"
)

func Templatize(s string, connString string, vars interface{}) (map[string]string, error) {
	var content map[string]string
	var err error
	if s == "" || strings.EqualFold(s, "git") {
		content, err = source.NewGitSource(connString).GetTemplateContent()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsupported source '%s'", s)
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
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, variables); err != nil {
		return "", fmt.Errorf("error processing template: %v", err)
	}
	return buf.String(), nil
}
