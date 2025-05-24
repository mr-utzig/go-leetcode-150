package rotate_array

import "testing"

func TestAlgorithm(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Algorithm(tt.args.nums, tt.args.k)
		})
	}
}
