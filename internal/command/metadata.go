package command

import (
	"fmt"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

var metadataFlags = []external.Flag{
	{
		Name:    flagInit,
		Type:    "boolean",
		Default: "false",
		Usage:   "true if the action is 'init', otherwise else",
	},
}

func runMetadata(flags *pflag.FlagSet) (plugin.SubcommandMetadata, error) {
	initFlag, _ := flags.GetBool(flagInit)
	if initFlag {
		return initMetadata, nil
	}
	return plugin.SubcommandMetadata{}, fmt.Errorf("unrecognized flag")
}
