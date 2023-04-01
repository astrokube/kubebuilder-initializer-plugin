package info

import "testing"

func TestAuthor(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "simple test",
			want: "The Astrokube Team <developer@astrokube.com>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Author(); got != tt.want {
				t.Errorf("Author() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "check hardcoded value",
			want: "v1alpha",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Version(); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}
