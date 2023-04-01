package command

import (
	"fmt"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

var flagsFlags = []external.Flag{
	{
		Name:    flagInit,
		Type:    "boolean",
		Default: "false",
		Usage:   "true if the action is 'init', otherwise else",
	},
}

func runFlags(flags *pflag.FlagSet) ([]external.Flag, error) {
	initFlag, _ := flags.GetBool(flagInit)
	if initFlag {
		return initFlags, nil
	} else {
		return nil, fmt.Errorf("unrecognized flag")
	}
}
