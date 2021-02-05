package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))
	fmt.Println(stocks["MSFT"])
	fmt.Println(stocks["TSLA"])

	value, ok := stocks["TSLA"]

	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	for key := range stocks {
		fmt.Println(key)
	}

	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
