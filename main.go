package main

import "fmt"

const (
	usdToEurRate float64 = 0.85
	usdToRubRate float64 = 110.0
)

func convertCurrency(amount float64, inputCurrency string, outputCurrency string) float64 {
	if amount == 0 {
		return 0.0
	}

	switch {
	case inputCurrency == "usd" && outputCurrency == "eur":
		return amount * usdToEurRate
	case inputCurrency == "usd" && outputCurrency == "rub":
		return amount * usdToRubRate
	case inputCurrency == "eur" && outputCurrency == "usd":
		return amount / usdToEurRate
	case inputCurrency == "eur" && outputCurrency == "rub":
		return amount * usdToRubRate / usdToEurRate
	case inputCurrency == "rub" && outputCurrency == "usd":
		return amount / usdToRubRate
	case inputCurrency == "rub" && outputCurrency == "eur":
		return amount / (usdToRubRate * usdToEurRate)
	}

	return 0.0
}

func scanCurrency() string {
	fmt.Print("Введите исходную валюту (usd, eur, rub): ")
	var inputCurrency string
	fmt.Scan(&inputCurrency)
	if inputCurrency != "usd" && inputCurrency != "eur" && inputCurrency != "rub" {
		fmt.Println("Неверная валюта")
		return scanCurrency()
	}
	return inputCurrency
}

func scanOutputCurrency() string {
	fmt.Print("Введите целевую валюту (usd, eur, rub): ")
	var outputCurrency string
	fmt.Scan(&outputCurrency)
	if outputCurrency != "usd" && outputCurrency != "eur" && outputCurrency != "rub" {
		fmt.Println("Неверная валюта")
		return scanOutputCurrency()
	}
	return outputCurrency
}

func scanAmount() float64 {
	fmt.Print("Введите сумму: ")
	var amount float64
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Сумма должна быть положительной")
		return scanAmount()
	}
	return amount
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")

	for {
		inputCurrency := scanCurrency()
		outputCurrency := scanOutputCurrency()
		amount := scanAmount()

		fmt.Println("result", convertCurrency(amount, inputCurrency, outputCurrency))
	}
}
