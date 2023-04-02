package templatizer

import (
	"reflect"
	"testing"
)

func Test_readVarsFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty yaml file",
			args:    args{path: "testdata/empty.yaml"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "missing yaml file",
			args:    args{path: "testdata/not-existing.yaml"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid yaml file",
			args: args{path: "testdata/vars.yml"},
			want: map[string]interface{}{
				"parent": map[string]interface{}{
					"child1": "value",
					"child2": map[string]interface{}{
						"child21": 25,
					},
				},
				"siblings": []interface{}{
					map[string]interface{}{
						"firstName": "David",
						"lastName":  "Doe",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readVarsFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readVarsFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readVarsFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
