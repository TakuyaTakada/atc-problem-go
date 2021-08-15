package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		s, t int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{1, 0},
			4,
		},
		{
			args{2, 5},
			10,
		},
		{
			args{10, 10},
			213,
		},
		{
			args{30, 100},
			2471,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
