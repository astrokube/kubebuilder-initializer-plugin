package command

import (
	"github.com/astrokube/layout-kubebuilder-plugin/pkg/templatizer"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

var initFlags = []external.Flag{
	{
		Name:    flagLicense,
		Type:    "string",
		Default: "apache2",
		Usage:   "license to use to boilerplate, may be one of 'apache2', 'none'",
	},
	{
		Name:    flagOwner,
		Type:    "string",
		Default: "",
		Usage:   "owner to add to the copyright",
	},
	{
		Name:    flagRepo,
		Type:    "string",
		Default: "",
		Usage:   "repository name (e.g., github.com/user/repo).",
	},
	{
		Name:    "params",
		Type:    "array",
		Default: "",
		Usage:   "repository name (e.g., github.com/user/repo).",
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
	$ kubebuilder init --plugins astrokube-layout/v1 --layout github.com/astronetes/operator-template
	`,
}

func runInit(flags *pflag.FlagSet) (map[string]string, error) {
	/**
	license, _ := flags.GetString(flagLicense)
	owner, _ := flags.GetString(flagOwner)
	repo, _ := flags.GetString(flagRepo)


	templater.TemplateFromRepository(from)
	**/
	source, _ := flags.GetString(flagSource)
	from, _ := flags.GetString(flagFrom)
	return templatizer.Templatize(source, from, nil)
}
