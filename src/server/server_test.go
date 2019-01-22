package server

import (
	"testing"
)

func TestRequestStockPrices(t *testing.T) {
	var data InputParams
	data.Symbols = []string{"AAPL", "MSFT"}
	data.Stock_exchange = []string{"NYSE", "NASDAQ"}
	_, err := data.RequestStockPrices()
	if err != nil {
		t.Errorf("Failed")
	}

}
