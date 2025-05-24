package jump_game

import (
	"testing"
)

func TestAlgorithm(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 0, 1, 0}},
			want: false,
		},
		{
			name: "Example 4",
			args: args{nums: []int{1, 2}},
			want: true,
		},
		{
			name: "Example 5",
			args: args{nums: []int{2, 0}},
			want: true,
		},
		{
			name: "Example 6",
			args: args{nums: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "Example 7",
			args: args{nums: []int{1, 1, 1, 0}},
			want: true,
		},
		{
			name: "Example 8",
			args: args{nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}},
			want: true,
		},
		{
			name: "Example 9",
			args: args{nums: []int{2, 0, 0, 0, 2, 0, 0, 0}},
			want: false,
		},
		{
			name: "Example 10",
			args: args{nums: []int{4, 1, 0, 2, 2, 4, 4, 4, 1, 2}},
			want: true,
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
