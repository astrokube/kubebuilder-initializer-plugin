package variables

import (
	"bytes"
	"fmt"
	"text/template"
)

func ReplaceVariables(name string, content string, variables map[string]interface{}) (string, error) {
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
