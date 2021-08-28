package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n, q int
		a, k []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{4, 3, []int{3,5,6,7}, []int{2,5,3}},
			[]int{2,9,4},
		},
		{
			args{5, 2, []int{1,2,3,4,5}, []int{1,10}},
			[]int{6,15},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.q, tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
