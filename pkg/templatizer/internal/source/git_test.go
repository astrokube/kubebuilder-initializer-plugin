package source

import (
	"reflect"
	"testing"
)

func TestNewGitSource(t *testing.T) {
	type args struct {
		conn string
	}
	tests := []struct {
		name string
		args args
		want *GitSource
	}{
		{
			name: "Basic test",
			args: args{conn: "https://github.com/astrokube/kubebuilder-operator-template"},
			want: &GitSource{
				url:       "https://github.com/astrokube/kubebuilder-operator-template",
				refOrigin: "origin",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitSource(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
