package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"0601889"},
			"6881090",
		},
		{
			args{"86910"},
			"01698",
		},
		{
			args{"01010"},
			"01010",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
