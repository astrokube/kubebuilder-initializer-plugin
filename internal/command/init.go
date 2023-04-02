package command

import (
	"github.com/astrokube/layout-kubebuilder-plugin/pkg/templatizer"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

var initFlags = []external.Flag{
	{
		Name:    "vars",
		Type:    "string",
		Default: ".kubebuilder-layout.yaml",
		Usage:   "path to the file that contains the variables to be used. By default the plugin uses the file in path .kubebuilder-layout.yaml",
	},
	{
		Name:    flagFrom,
		Type:    "string",
		Default: "",
		Usage: "repository path (e.g., github.com/my-organization/my-repo). We can pass a specific branch " +
			"after the repository path (e.g., github.com/my-organization/my-repo#develop) or the credentials If " +
			"required (e.g., username:password@github.com/my-organization/my-repo#develop)",
	},
}

var initMetadata = plugin.SubcommandMetadata{
	Description: "The `init` subcommand of the layout plugin is meant to initialize a project via Kubebuilder." +
		"It scaffolds a repository with a previously created layout",
	Examples: `
	Scaffold with the defaults:
	$ kubebuilder init --plugins astrokube-layout/v1

	Scaffold with a specific layout:
	$ kubebuilder init --plugins astrokube-layout/v1 --layout github.com/astronetes/operator-template --vars custom-vars.yml
	`,
}

func runInit(flags *pflag.FlagSet) (map[string]string, error) {

	source, _ := flags.GetString(flagSource)
	from, _ := flags.GetString(flagFrom)
	vars, _ := flags.GetString(flagVars)
	return templatizer.Templatize(source, from, vars)
}
