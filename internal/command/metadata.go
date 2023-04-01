package command

import (
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

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

type metadataCmd struct {
}

func (cmd *metadataCmd) name() string {
	return "metadata"
}

func (cmd *metadataCmd) flags() []external.Flag {
	return []external.Flag{
		{
			Name:    flagInit,
			Type:    "boolean",
			Default: "false",
			Usage:   "true if the action is 'init', otherwise else",
		},
	}
}

func (cmd *metadataCmd) run(flags *pflag.FlagSet, response *external.PluginResponse) {
	initFlag, _ := flags.GetBool(flagInit)
	if initFlag {
		response.Metadata = initMetadata
	} else {
		response.Error = true
		response.ErrorMsgs = []string{
			"unrecognized flag",
		}
	}
}
