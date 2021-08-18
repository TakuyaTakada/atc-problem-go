package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		a, b, c int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{2, 5, 2},
			5,
		},
		{
			args{4, 5, 6},
			0,
		},
		{
			args{1, 1, 1},
			1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
