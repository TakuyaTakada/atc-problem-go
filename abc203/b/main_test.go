package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n, k int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{1,2},
			203,
		},
		{
			args{3,3},
			1818,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
