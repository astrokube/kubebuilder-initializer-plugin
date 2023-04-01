package command

import (
	"fmt"

	"github.com/astrokube/layout-kubebuilder-plugin/pkg/info"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

//type Command func(*external.PluginRequest) external.PluginResponse

type Command interface {
	name() string
	flags() []external.Flag
	run(*pflag.FlagSet, *external.PluginResponse)
}

type cmdKey string

const (
	flagLicense = "license"
	flagOwner   = "owner"
	flagRepo    = "repo"
	flagFrom    = "from"
	flagSource  = "source"
	flagInit    = "init"
)

const (
	ActionFlags    cmdKey = "flags"
	ActionInit     cmdKey = "init"
	ActionMetadata cmdKey = "metadata"
)

var cmdMap = map[cmdKey]Command{
	ActionFlags:    &flagsCmd{},
	ActionMetadata: &metadataCmd{},
	ActionInit:     &initCmd{},
}

func Run(request *external.PluginRequest) external.PluginResponse {
	cmd, ok := cmdMap[cmdKey(request.Command)]
	if !ok {
		return external.PluginResponse{
			Error: true,
			ErrorMsgs: []string{
				"unknown subcommand:" + request.Command,
			},
		}
	}
	response := &external.PluginResponse{
		APIVersion: info.Version(),
		Command:    cmd.name(),
		Universe:   request.Universe,
	}
	flags := processFlags(cmd, request.Args)
	cmd.run(flags, response)
	return *response
}

func processFlags(cmd Command, args []string) *pflag.FlagSet {
	flags := pflag.NewFlagSet(fmt.Sprintf("%sFlags", cmd.name()), pflag.ContinueOnError)
	for _, f := range cmd.flags() {
		switch f.Type {
		case "string":
			flags.String(f.Name, f.Default, f.Usage)
		case "boolean":
			flags.Bool(f.Name, f.Default == "true", f.Usage)
		case "array":
			flags.StringArray(f.Name, []string{}, f.Usage)
		default:
			flags.String(f.Name, f.Default, f.Usage)
		}

	}
	_ = flags.Parse(args)

	return flags
}
