package command

import (
	"os"
	"path"

	"github.com/astrokube/kubebuilder-initializer-plugin/pkg/templatizer"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

var initFlags = []external.Flag{
	{
		Name:    flagSource,
		Type:    "string",
		Default: "git",
		Usage:   "Type of source. The only supported value is 'git'",
	},
	{
		Name:    flagVars,
		Type:    "string",
		Default: ".kubebuilder-initializer.yaml",
		Usage: "path to the file that contains the variables to be used. By default the plugin uses the file in path ." +
			"kubebuilder-initializer.yaml",
	},
	{
		Name:    flagFrom,
		Type:    "string",
		Default: "",
		Usage:   "repository path (e.g., github.com/my-organization/my-repo).",
	},
	/**
	{
		Name:    "component-config",
		Type:    "boolean",
		Default: "false",
		Usage:   "create a versioned ComponentConfig file, may be 'true' or 'false'",
	},
	{
		Name:    "domain",
		Type:    "string",
		Default: "my.domain",
		Usage:   "domain for groups (default \"my.domain\")",
	},
	{
		Name:    "fetch-deps",
		Type:    "boolean",
		Default: "true",
		Usage:   "ensure dependencies are downloaded (default true)",
	},
	{
		Name:    "license",
		Type:    "string",
		Default: "",
		Usage:   "license to use to boilerplate, may be one of 'apache2', 'none' (default \"apache2\")",
	},
	{
		Name:    "owner",
		Type:    "string",
		Default: "",
		Usage:   "owner to add to the copyright",
	},
	**/

	/**
	  --project-name string      name of this project
	  --project-version string   project version (default "3")
	  --repo string              name to use for go module (e.g., github.com/user/repo), defaults to the go package of the current working directory.
	  --skip-go-version-check    if specified, skip checking the Go versio
	*/
}

var initMetadata = plugin.SubcommandMetadata{
	Description: "The `init` subcommand of the layout plugin is meant to initialize a project via Kubebuilder." +
		"It scaffolds a repository with a previously created layout",
	Examples: `
	Scaffold with the defaults:
	$ kubebuilder init --plugins kubebuilder-initializer-plugin/v1-alpha \
		--from https://github.com/astronetes/operator-template

	Scaffold overwriting the default file for the variables:
	$ kubebuilder init --plugins kubebuilder-initializer-plugin/v1-alpha \
		--from https://github.com/astronetes/operator-template \
		--vars path-to-my-custom-variables.yml
		
	`,
}

func runInit(flags *pflag.FlagSet) (map[string]string, error) {
	source, err := flags.GetString(flagSource)
	if err != nil {
		return nil, err
	}
	from, err := flags.GetString(flagFrom)
	if err != nil {
		return nil, err
	}
	vars, err := flags.GetString(flagVars)
	if err != nil {
		return nil, err
	}
	content, err := templatizer.Templatize(source, from, vars)
	if err != nil {
		return content, err
	}
	for k := range content {
		parentDir := path.Dir(k)[1:]
		if parentDir != "" {
			if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	return content, nil
}
