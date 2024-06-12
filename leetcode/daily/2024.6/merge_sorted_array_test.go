package daily

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
				want:  []int{1, 2, 2, 3, 5, 6},
			},
		},
		{
			name: "test case 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
				want:  []int{1},
			},
		},
		{
			name: "test case 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
				want:  []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)

			//	want == tt.args.nums1
			if len(tt.args.nums1) != len(tt.args.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.args.want)
			}

			for i := 0; i < len(tt.args.nums1); i++ {
				if tt.args.nums1[i] != tt.args.want[i] {
					t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.args.want)
					break
				}
			}
		})
	}
}
