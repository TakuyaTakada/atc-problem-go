package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{[]int{5,1,3}},
			"Yes",
		},
		{
			args{[]int{1,4,3}},
			"No",
		},
		{
			args{[]int{5,5,5}},
			"Yes",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.a); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
