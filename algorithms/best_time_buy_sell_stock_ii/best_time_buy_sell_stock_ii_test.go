package best_time_buy_sell_stock_ii

import "testing"

func TestPrices1(t *testing.T) {
	res := Algorithm(Prices1)

	if res != 7 {
		t.Errorf(`Algorithm(%v) = %v, expected 7`, Prices1, res)
	}
}

func TestPrices2(t *testing.T) {
	res := Algorithm(Prices2)

	if res != 4 {
		t.Errorf(`Algorithm(%v) = %v, expected 4`, Prices2, res)
	}

}

func TestPrices3(t *testing.T) {
	res := Algorithm(Prices3)

	if res != 0 {
		t.Errorf(`Algorithm(%v) = %v, expected 0`, Prices3, res)
	}

}
