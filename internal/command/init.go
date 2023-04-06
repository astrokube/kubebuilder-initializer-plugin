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
		Name:    "vars",
		Type:    "string",
		Default: ".kubebuilder-initializer.yaml",
		Usage: "path to the file that contains the variables to be used. By default the plugin uses the file in path ." +
			"kubebuilder-initializer.yaml",
	},
	{
		Name:    flagFrom,
		Type:    "string",
		Default: "",
		Usage: "repository path (e.g., github.com/my-organization/my-repo). We can pass a specific branch " +
			"after the repository path (e.g., github.com/my-organization/my-repo#develop) or the credentials If " +
			"required (e.g., username:password@github.com/my-organization/my-repo#develop)",
	},
	/**
	{
		Name:    flagDomain,
		Type:    "string",
		Default: "",
		Usage:   "store the domain of the project. ",
	},
	{
		Name:    flagRepo,
		Type:    "string",
		Default: "",
		Usage: "name to use for go module (e.g., github.com/user/repo), defaults to the go package of the current " +
			"working directory.",
	},
	{
		Name:    flagProjectName,
		Type:    "string",
		Default: "",
		Usage:   "the name of the project. This will be used to scaffold the manager data",
	},
	*/
}

var initMetadata = plugin.SubcommandMetadata{
	Description: "The `init` subcommand of the layout plugin is meant to initialize a project via Kubebuilder." +
		"It scaffolds a repository with a previously created layout",
	Examples: `
	Scaffold with the defaults:
	$ kubebuilder init --plugins astrokube-layout/v1

	Scaffold with a specific layout:
	$ kubebuilder init --plugins astrokube-layout/v1 --layout github.com/astronetes/operator-template 
		--vars custom-vars.yml
	`,
}

func runInit(flags *pflag.FlagSet) (map[string]string, error) {
	source, _ := flags.GetString(flagSource)
	from, _ := flags.GetString(flagFrom)
	vars, _ := flags.GetString(flagVars)
	/**
	domain, _ := flags.GetString(flagDomain)
	projectName, _ := flags.GetString(flagProjectName)
	repo, _ := flags.GetString(flagRepo)
	**/
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
