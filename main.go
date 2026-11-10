package main

import "fmt"

func main() {
	hello_convertor()
	fromCurrency()
	output_currency()
}
func hello_convertor() {
	fmt.Println("Добро пожаловать в конвертор валют!")
}
func fromCurrency() string {
	for {
		curencyUser := " "
		fmt.Printf("Доступные исходные валюты:\n1)USD\n2)EUR\n3)CNY\n4)JPY\n")
		fmt.Printf("Ввод исходной валюты: ")
		fmt.Scan(&curencyUser)
		switch curencyUser {
		case "USD", "1":
			fmt.Println("Вы выбрали: USD")
			return curencyUser
		case "EUR", "2":
			fmt.Println("Вы выбрали: EUR")
			return curencyUser
		case "CNY", "3":
			fmt.Println("Вы выбрали: CNY")
			return curencyUser
		case "JPY", "4":
			fmt.Println("Вы выбрали: JPY")
			return curencyUser
		default:
			fmt.Println("Такой валюты нет!")
			continue
		}

	}
}
func selectedCurrency() string {
	for {
		fmt.Println("Выберите вылюту в которую хотите перевести")
		fmt.Printf("Доступные выходные валюты:\n1)RUB\n2)GBP\n3)KZT\n4)TRY\n")
		outCurrency := ""
		fmt.Scan(&outCurrency)
		switch outCurrency {
		case "RUB", "1":
			fmt.Println("Вы выбрали: RUB")
			return outCurrency
		case "GBP", "2":
			fmt.Println("Вы выбрали: GBP")
			return outCurrency
		case "KZT", "3":
			fmt.Println("Вы выбрали: KZT")
			return outCurrency
		case "TRY", "4":
			fmt.Println("Вы выбрали: TRY")
			return outCurrency
		default:
			fmt.Println("Такой валюты нет!")
			continue
		}
	}

}
func sumCurrency() int {
	fmt.Println("Ввведите сумму для обмена:")
	sum := 0
	fmt.Scan(&sum)
	return sum
}
func currencyConverter() float64 {
	sum := sumCurrency()
	result := 0.0
	inputUserCurrency := fromCurrency()
	outoutUserCurrency := selectedCurrency()
	if inputUserCurrency == "USD" && outoutUserCurrency == "RUB" {
		result = float64(sum) * 83.01
	} else if inputUserCurrency == "USD" && outoutUserCurrency == "GBP" {
		result = float64(sum) * 0.78

	} else if inputUserCurrency == "USD" && outoutUserCurrency == "KZT" {
		result = float64(sum) * 447.14

	} else if inputUserCurrency == "USD" && outoutUserCurrency == "TRY" {

		result = float64(sum) * 34.20
	}
	if inputUserCurrency == "EUR" && outoutUserCurrency == "RUB" {
		result = float64(sum) * 90.11
	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "GBP" {
		result = float64(sum) * 0.85

	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "KZT" {
		result = float64(sum) * 485.20

	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "TRY" {

		result = float64(sum) * 37.10
	}
	if inputUserCurrency == "CNY" && outoutUserCurrency == "RUB" {
		result = float64(sum) * 11.49
	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "GBP" {
		result = float64(sum) * 0.11

	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "KZT" {
		result = float64(sum) * 61.40

	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "TRY" {

		result = float64(sum) * 4.70
	}
	if inputUserCurrency == "JPY" && outoutUserCurrency == "RUB" {
		result = float64(sum) * 0.53
	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "GBP" {
		result = float64(sum) * 0.0051

	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "KZT" {
		result = float64(sum) * 2.85

	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "TRY" {

		result = float64(sum) * 0.22
	}

	return result
}

func output_currency() {
	res := currencyConverter()
	inputCurrency := fromCurrency()
	outputCurrency := selectedCurrency()
	fmt.Printf("Вы выбрали обменять: %s на %s\nИтоговая сумма которую вы получаете: %.2f", inputCurrency, outputCurrency, res)
}
