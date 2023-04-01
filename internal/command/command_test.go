package command

import (
	"reflect"
	"testing"

	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/external"
)

func TestRun(t *testing.T) {
	type args struct {
		request *external.PluginRequest
	}
	tests := []struct {
		name string
		args args
		want external.PluginResponse
	}{
		{
			name: "unknown command",
			args: args{request: &external.PluginRequest{
				Command: "invalidCommand",
			}},
			want: external.PluginResponse{
				Error:     true,
				ErrorMsgs: []string{"unknown command 'invalidCommand'"},
			},
		},
		{
			name: "missing command",
			args: args{request: &external.PluginRequest{}},
			want: external.PluginResponse{
				Error:     true,
				ErrorMsgs: []string{"missing command"},
			},
		},
		{
			name: "metadata command, missing required valid flag",
			args: args{request: &external.PluginRequest{
				Args:    []string{"--unknown"},
				Command: "metadata",
			}},
			want: external.PluginResponse{
				Error:     true,
				ErrorMsgs: []string{"unrecognized flag"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
