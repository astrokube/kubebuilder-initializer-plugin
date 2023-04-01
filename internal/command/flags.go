package command

import (
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

type flagsCmd struct {
}

func (cmd *flagsCmd) name() string {
	return "flags"
}

func (cmd *flagsCmd) flags() []external.Flag {
	return []external.Flag{
		{
			Name:    flagInit,
			Type:    "boolean",
			Default: "false",
			Usage:   "true if the action is 'init', otherwise else",
		},
	}
}

func (cmd *flagsCmd) run(flags *pflag.FlagSet, response *external.PluginResponse) {
	initFlag, _ := flags.GetBool(flagInit)
	if initFlag {
		response.Flags = cmdMap[ActionInit].flags()
	} else {
		response.Error = true
		response.ErrorMsgs = []string{
			"unrecognized flag",
		}
	}
}
