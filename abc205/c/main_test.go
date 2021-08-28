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
		want string
	}{
		{
			args{3,2,4},
			">",
		},
		{
			args{-7,7,2},
			"=",
		},
		{
			args{-8,6,3},
			"<",
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
