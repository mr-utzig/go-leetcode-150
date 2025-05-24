package majority_element

import "testing"

func TestAlgorithm(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Algorithm(tt.args.nums); got != tt.want {
				t.Errorf("Algorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
