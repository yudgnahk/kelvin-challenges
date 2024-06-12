package daily

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "test case 2",
			args: args{nums: []int{-1, -100, 3, 99}, k: 2},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "test case 3",
			args: args{nums: []int{1, 2, 3}, k: 4},
			want: []int{3, 1, 2},
		},
		{
			name: "test case 4",
			args: args{nums: []int{1, 2, 3}, k: 0},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			for i := 0; i < len(tt.want); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
				}
			}
		})
	}
}
