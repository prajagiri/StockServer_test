package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	url       = "https://www.worldtradingdata.com/api/v1/stock?symbol="
	api_token = "5LM183iQ3mKUWX8s6txwwfGX1Qtz0WU7ekLo6zLf1LJM9gIx9MWJ6PI5S06Z"
)

type InputParams struct {
	Symbols        []string
	Stock_exchange []string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var data InputParams
	for key, value := range r.Form {
		if key == "symbol" {
			data.Symbols = value
		} else if key == "stock_exchange" {
			data.Stock_exchange = value
		} else {
			fmt.Fprintf(w, "Invalid url - %s", r.URL)
			return
		}
	}
	res, err := data.RequestStockPrices()
	if err != nil {
		fmt.Fprintf(w, "%s - %s", res, err)
		return
	}
	fmt.Fprintf(w, "%s", res)
	return
}

func (data *InputParams) RequestStockPrices() (string, error) {
	var formURL string
	if len(data.Stock_exchange) > 0 {
		formURL = url + strings.Join(data.Symbols, ",") + "&stock_exchange=" + strings.Join(data.Stock_exchange, ",") + "&api_token=" + api_token
	} else {
		formURL = url + strings.Join(data.Symbols, ",") + "&stock_exchange=AMEX" + "&api_token=" + api_token
	}

	fmt.Println("Test - ", formURL)
	resp, err := http.Get(formURL)
	if err != nil {
		log.Fatalln(err)
		return "Failed to get stock prices", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "Failed to read response", err
	}

	return string(body), nil
}

//https://www.worldtradingdata.com/api/v1/stock?symbol=AAPL,MSFT,HSBA.L&api_token=5LM183iQ3mKUWX8s6txwwfGX1Qtz0WU7ekLo6zLf1LJM9gIx9MWJ6PI5S06Z
