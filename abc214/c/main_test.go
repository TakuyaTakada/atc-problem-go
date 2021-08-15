package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n int
		s, t []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{3, []int{4,1,5}, []int{3,10,100}},
			[]int{3,7,8},
		},
		{
			args{4, []int{100,100,100,100}, []int{1,1,1,1}},
			[]int{1,1,1,1},
		},
		{
			args{4, []int{1,2,3,4}, []int{1,2,4,7}},
			[]int{1,2,4,7},
		},
		{
			args{8, []int{84,87,78,16,94,36,87,93}, []int{50,22,63,28,91,60,64,27}},
			[]int{50,22,63,28,44,60,64,27},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.s, tt.args.t); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
