package command

import (
	"fmt"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

const (
	flagVars = "vars"

	flagDomain      = "domain"
	flagRepo        = "repo"
	flagProjectName = "project-name"
	flagFrom        = "from"
	flagSource      = "source"
	flagInit        = "init"
)

const (
	ActionFlags         = "flags"
	ActionInit          = "init"
	ActionMetadata      = "metadata"
	ActionCreateAPI     = "create api"
	ActionCreateWebhook = "create webhook"
)

func Run(request *external.PluginRequest) external.PluginResponse {
	var response = external.PluginResponse{
		APIVersion: request.APIVersion,
		Command:    request.Command,
	}

	var err error

	switch request.Command {
	case ActionInit:
		flagSet := processFlags(request, initFlags)
		response.Universe, err = runInit(flagSet)
	case ActionFlags:
		flagSet := processFlags(request, flagsFlags)
		response.Flags, err = runFlags(flagSet)
	case ActionMetadata:
		flagSet := processFlags(request, metadataFlags)
		response.Metadata, err = runMetadata(flagSet)
	case ActionCreateAPI:
		break
	case ActionCreateWebhook:
		break
	case "":
		err = fmt.Errorf("missing command")
	default:
		err = fmt.Errorf("unknown command '%s'", request.Command)
	}

	if err != nil {
		response.Error = true
		response.ErrorMsgs = []string{err.Error()}
	}
	return response
}

func processFlags(request *external.PluginRequest, flags []external.Flag) *pflag.FlagSet {
	flagsSet := pflag.NewFlagSet(fmt.Sprintf("%sFlags", request.Command), pflag.ContinueOnError)
	for _, f := range flags {
		switch f.Type {
		case "string":
			flagsSet.String(f.Name, f.Default, f.Usage)
		case "boolean":
			flagsSet.Bool(f.Name, f.Default == "true", f.Usage)
		case "array":
			flagsSet.StringArray(f.Name, []string{}, f.Usage)
		default:
			flagsSet.String(f.Name, f.Default, f.Usage)
		}
	}
	_ = flagsSet.Parse(request.Args)

	return flagsSet
}
