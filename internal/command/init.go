package command

import (
	"github.com/astrokube/layout-kubebuilder-plugin/pkg/templatizer"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

type initCmd struct {
}

func (cmd *initCmd) name() string {
	return "init"
}

func (cmd *initCmd) flags() []external.Flag {
	return []external.Flag{
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
}

func (cmd *initCmd) run(flags *pflag.FlagSet, response *external.PluginResponse) {
	/**
	license, _ := flags.GetString(flagLicense)
	owner, _ := flags.GetString(flagOwner)
	repo, _ := flags.GetString(flagRepo)


	templater.TemplateFromRepository(from)
	**/
	println("---")
	flags.Visit(func(flag *pflag.Flag) {
		println(flag.Name)
		println(flag.Value.String())
	})
	source, _ := flags.GetString(flagSource)
	from, _ := flags.GetString(flagFrom)
	files, err := templatizer.Templatize(source, from, nil)
	if err != nil {
		response.Error = true
		response.ErrorMsgs = []string{
			err.Error(),
		}
		return
	}
	for path, content := range files {
		response.Universe[path] = content
	}
}
