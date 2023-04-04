package templatizer

import (
	"fmt"
	"os"
	"strings"

	"github.com/astrokube/kubebuilder-layout-plugin/pkg/templatizer/internal/source"
	"github.com/astrokube/kubebuilder-layout-plugin/pkg/templatizer/internal/variables"
	"gopkg.in/yaml.v3"
)

func Templatize(sourceType string, connString string, varsFile string) (map[string]string, error) {
	var content map[string]string
	var err error
	if sourceType == "" || strings.EqualFold(sourceType, "git") {
		content, err = source.NewGitSource(connString).GetTemplateContent()
		if err != nil {
			return nil, err
		}

		vars, err := readVarsFile(varsFile)
		if err != nil {
			return nil, err
		}
		procContent := make(map[string]string, len(content))
		for k, v := range content {
			processedFile, err := variables.ReplaceVariables(k, v, vars)
			if err != nil {
				return nil, err
			}
			procK, err := variables.ReplaceVariables(k, k, vars)
			if err != nil {
				return nil, err
			}
			procContent[procK] = processedFile
		}

		return procContent, nil
	}
	return nil, fmt.Errorf("unsupported source '%s'", sourceType)
}

func readVarsFile(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file with variables: '%w'", err)
	}
	out := make(map[string]interface{}, 0)
	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("error unmarshaling file with variables: '%w'", err)
	}

	return out, nil
}
