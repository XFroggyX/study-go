package main

import "fmt"

func main() {
	const usdToEurRate float64 = 0.85
	const usdToRubRate float64 = 110.0

	eurToRubRate := usdToRubRate / usdToEurRate

	fmt.Println(eurToRubRate)
}
