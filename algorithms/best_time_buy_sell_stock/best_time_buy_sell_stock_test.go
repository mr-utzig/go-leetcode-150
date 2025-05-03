package best_time_buy_sell_stock

import "testing"

func TestPrices1(t *testing.T) {
	expected := 5
	res := Algorithm(Prices1)

	if res != expected {
		t.Errorf(`Algorithm(%v) = %q, expected %v`, Prices1, res, expected)
	}
}

func TestPrices2(t *testing.T) {
	expected := 0
	res := Algorithm(Prices2)

	if res != expected {
		t.Errorf(`Algorithm(%v) = %q, expected %v`, Prices2, res, expected)
	}
}
