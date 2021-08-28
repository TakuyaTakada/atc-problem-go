package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		a, b int
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args{45, 200},
			90,
		},
		{
			args{37, 450},
		166.5,
		},
		{
			args{50, 0},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
