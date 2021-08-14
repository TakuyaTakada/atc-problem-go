package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		mt Mountains
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{Mountains{
				&Mount{"Everest", 8849},
				&Mount{"K2", 8611},
				&Mount{"Kangchenjunga", 8586},
			}},
			"K2",
		},
		{
			args{Mountains{
				&Mount{"Kita", 3193},
				&Mount{"Aino", 3189},
				&Mount{"Fuji", 3776},
				&Mount{"Okuhotaka", 3190},
			}},
			"Kita",
		},
		{
			args{Mountains{
				&Mount{"QCFium", 2846},
				&Mount{"chokudai", 2992},
				&Mount{"kyoprofriends", 2432},
				&Mount{"penguinman", 2390},
			}},
			"QCFium",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.mt); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
