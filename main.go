package main

import "fmt"

func scanCurrency() float64 {
	var currency float64
	fmt.Print("Enter currency: ")
	fmt.Scan(&currency)
	return currency
}

func calcCurrency(currency float64, inputCurrency string, outputCurrency string) float64 {
	return 0.0
}

func main() {
	const usdToEurRate float64 = 0.85
	const usdToRubRate float64 = 110.0

	eurToRubRate := usdToRubRate / usdToEurRate

	fmt.Println(eurToRubRate)
}
