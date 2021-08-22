package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"aab", 2},
			"aba",
		},
		{
			args{"baba", 4},
			"baab",
		},
		{
			args{"ydxwacbz", 40320},
			"zyxwdcba",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
