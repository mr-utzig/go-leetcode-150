package merge_sorted_array

import "testing"

func TestAlgorithm(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1},
				m:     3,
				nums2: []int{},
				n:     0,
			},
		},
		{
			name: "Example 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Algorithm(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}
