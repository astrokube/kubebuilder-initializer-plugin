package command

import (
	"testing"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

func Test_initCmd_run(t *testing.T) {
	type args struct {
		flags    *pflag.FlagSet
		response *external.PluginResponse
	}

	flags := processFlags(&initCmd{}, []string{
		"--from=github.com/astrokube/acm-examples",
		"-- --a=john --b=dave",
	})

	tests := []struct {
		name string
		cmd  *initCmd
		args args
	}{
		{
			name: "from basic repository",
			cmd:  &initCmd{},
			args: args{
				flags: flags,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &initCmd{}
			cmd.run(tt.args.flags, tt.args.response)
		})
	}
}
