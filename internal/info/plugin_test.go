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
