package utils

import "testing"

func TestGetRandNum(t *testing.T) {
	type args struct {
		min int
		max int
	}
	for i := 0; i < 100; i++ {
		tests := []struct {
			name string
			args args
			want bool
		}{
			{
				name: "TestGetRandNum",
				args: struct {
					min int
					max int
				}{min: i, max: i + 100},
				want: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := GetRandNum(tt.args.min, tt.args.max); !(got >= tt.args.min && got < tt.args.max) {
					t.Errorf("GetRandNum() = %v,min = %d, max=%d, want %v",
						got, tt.args.min, tt.args.max, tt.want)
				}
			})
		}

	}
}
