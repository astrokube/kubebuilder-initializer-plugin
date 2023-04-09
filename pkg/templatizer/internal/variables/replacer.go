// Package variables is used to deal with variables and the replacement of values in Go templates
package variables

import (
	"bytes"
	"fmt"
	"text/template"
)

// ReplaceVariables is a function that create a Go text template from the given content and replace the provided
// variables in the map.
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
