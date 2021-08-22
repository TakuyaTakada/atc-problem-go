package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		s Chars
	}
	tests := []struct {
		args args
		want Chars
	}{
		{
			args{
				Chars{"a", "a", "b"},
			},
			Chars{"aba", "baa"},
		},
		{
			args{
				Chars{"a", "b", "c"},
			},
			Chars{"acb", "bac", "bca", "cab", "cba"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			n := tt.want.Len()
			for j := 0; j < n; j++ {
				if got := nextPermutation(tt.args.s); got != true {
					t.Errorf("nextPermutation() = %v, want %v", got, true)
				}
				if str := strings.Join(tt.args.s, ""); str != tt.want[j] {
					t.Errorf("join = %v, want %v", str, tt.want[j])
				}
			}
			if got := nextPermutation(tt.args.s); got != false {
				t.Errorf("nextPermutation() = %v, want %v", got, false)
			}
		})
	}
}