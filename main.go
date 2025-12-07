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

	if inputCurrency == "usd" && outputCurrency == "eur" {
		return amount * usdToEurRate
	} else if inputCurrency == "usd" && outputCurrency == "rub" {
		return amount * usdToRubRate
	} else if inputCurrency == "eur" && outputCurrency == "usd" {
		return amount / usdToEurRate
	} else if inputCurrency == "eur" && outputCurrency == "rub" {
		return amount * usdToRubRate / usdToEurRate
	} else if inputCurrency == "rub" && outputCurrency == "usd" {
		return amount / usdToRubRate
	} else if inputCurrency == "rub" && outputCurrency == "eur" {
		return amount / (usdToRubRate * usdToEurRate)
	}

	return 0.0
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")

	for {
		fmt.Print("Введите исходную валюту (usd, eur, rub): ")
		var inputCurrency string
		fmt.Scan(&inputCurrency)
		if inputCurrency != "usd" && inputCurrency != "eur" && inputCurrency != "rub" {
			fmt.Println("Неверная валюта")
			continue
		}
		fmt.Print("Введите целевую валюту (usd, eur, rub): ")
		var outputCurrency string
		fmt.Scan(&outputCurrency)
		if outputCurrency != "usd" && outputCurrency != "eur" && outputCurrency != "rub" {
			fmt.Println("Неверная валюта")
			continue
		} else if inputCurrency == outputCurrency {
			fmt.Println("Исходная и целевая валюты не могут быть одинаковыми")
			continue
		}
		fmt.Print("Введите сумму: ")
		var amount float64
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Println("Сумма должна быть положительной")
			continue
		}

		fmt.Println(convertCurrency(amount, inputCurrency, outputCurrency))
	}
}
