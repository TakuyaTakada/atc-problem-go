package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"Hello,World!"},
			"AC",
		},
		{
			args{"Hello!World!"},
			"WA",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
