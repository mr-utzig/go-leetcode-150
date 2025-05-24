package best_time_buy_sell_stock_ii

import (
	"testing"
)

func TestAlgorithm(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Algorithm(tt.args.prices); got != tt.want {
				t.Errorf("Algorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
