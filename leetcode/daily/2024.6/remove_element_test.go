package daily

import (
	"slices"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name         string
		args         args
		expectedNums []int
		want         int
	}{
		{
			name:         "test case 1",
			args:         args{nums: []int{3, 2, 2, 3}, val: 3},
			expectedNums: []int{2, 2},
			want:         2,
		},
		{
			name:         "test case 2",
			args:         args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			expectedNums: []int{0, 0, 1, 3, 4},
			want:         5,
		},
		{
			name:         "test case 3",
			args:         args{nums: []int{1}, val: 1},
			expectedNums: []int{},
			want:         0,
		},
		{
			name:         "test case 4",
			args:         args{nums: []int{1}, val: 2},
			expectedNums: []int{1},
			want:         1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}

			slices.Sort(tt.args.nums)
			slices.Sort(tt.expectedNums)
			for i := 0; i < tt.want; i++ {
				if tt.args.nums[i] != tt.expectedNums[i] {
					t.Errorf("removeElement() = %v, want %v", tt.args.nums, tt.expectedNums)
				}
			}
		})
	}
}
