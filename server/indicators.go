package main

import (
	"fmt"
)

func getLastQuote(interval string, ticker string) float64 {
	response := ChartResponse{}
	error := getJson("https://query1.finance.yahoo.com/v7/finance/chart/"+ticker+"?range="+interval+"&interval="+interval+"&indicators=quote&includeTimestamps=true&includePrePost=true&corsDomain=finance.yahoo.com", &response)

	if error != nil {
		return 0
	}

	fmt.Println(response.Chart.Result[0].Indicators.Quote[0].Close)
	return -1
}
