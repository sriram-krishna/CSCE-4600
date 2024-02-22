package builtins_test

import (
	"testing"

	"github.com/sriram-krishna/CSCE-4600/Project2/builtins"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "no args",
			args: []string{},
			want: "",
		},
		{
			name: "one arg",
			args: []string{"hello"},
			want: "hello",
		},
		{
			name: "multiple args",
			args: []string{"hello", "world"},
			want: "hello world",
		},
		{
			name: "args with spaces",
			args: []string{"hello", "world", "with spaces"},
			want: "hello world with spaces",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builtins.Echo(tt.args...); got != tt.want {
				t.Errorf("Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}
