package command

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

const (
	flagVars   = "vars"
	flagFrom   = "from"
	flagSource = "source"
	flagInit   = "init"
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
	var flagSet *pflag.FlagSet

	switch request.Command {
	case ActionInit:
		flagSet, err = processFlags(request, initFlags)
		if err != nil {
			break
		}
		response.Universe, err = runInit(flagSet)
	case ActionFlags:
		flagSet, err = processFlags(request, flagsFlags)
		if err != nil {
			break
		}
		response.Flags, err = runFlags(flagSet)
	case ActionMetadata:
		flagSet, err = processFlags(request, metadataFlags)
		if err != nil {
			break
		}
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

func processFlags(request *external.PluginRequest, flags []external.Flag) (*pflag.FlagSet, error) {
	flagsSet := pflag.NewFlagSet(fmt.Sprintf("%sFlags", request.Command), pflag.ContinueOnError)
	for _, f := range flags {
		switch f.Type {
		case "string":
			flagsSet.String(f.Name, f.Default, f.Usage)
		case "boolean":
			flagsSet.Bool(f.Name, f.Default == "true", f.Usage)
		default:
			flagsSet.String(f.Name, f.Default, f.Usage)
		}
	}
	args := filterProvidedArgs(request.Args, flags)
	if err := flagsSet.Parse(args); err != nil {
		return nil, err
	}
	return flagsSet, nil
}

func filterProvidedArgs(args []string, supportedFlags []external.Flag) (out []string) {
	supportedFlagsMap := make(map[string]struct{}, len(supportedFlags))
	for i := range supportedFlags {
		key := fmt.Sprintf("--%s", supportedFlags[i].Name)
		supportedFlagsMap[key] = struct{}{}
	}
	expectArgValue := false
	for i := range args {
		arg := args[i]
		if expectArgValue {
			out = append(out, arg)
		} else if _, ok := supportedFlagsMap[args[i]]; ok {
			out = append(out, arg)
		}
		expectArgValue = strings.HasPrefix(arg, "--")
	}
	return
}
