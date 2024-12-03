package daily

import "testing"

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				s:      "LeetcodeHelpsMeLearn",
				spaces: []int{8, 13, 15},
			},
			want: "Leetcode Helps Me Learn",
		},
		{
			name: "test case 2",
			args: args{
				s:      "icodeingolang",
				spaces: []int{1, 5, 7, 9},
			},
			want: "i code in go lang",
		},
		{
			name: "test case 3",
			args: args{
				s:      "spacing",
				spaces: []int{0, 1, 2, 3, 4, 5, 6},
			},
			want: " s p a c i n g",
		},
		{
			name: "test case 4",
			args: args{
				s:      "p",
				spaces: []int{0},
			},
			want: " p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
