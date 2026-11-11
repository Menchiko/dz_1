package main

import (
	"fmt"
	"strconv"
)

func main() {
	hello_convertor()
	result := currencyConverter()
	fmt.Println(result)
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
			return "USD"
		case "EUR", "2":
			fmt.Println("Вы выбрали: EUR")
			return "EUR"
		case "CNY", "3":
			fmt.Println("Вы выбрали: CNY")
			return "CNY"
		case "JPY", "4":
			fmt.Println("Вы выбрали: JPY")
			return "JPY"
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
			return "RUB"
		case "GBP", "2":
			fmt.Println("Вы выбрали: GBP")
			return "GBP"
		case "KZT", "3":
			fmt.Println("Вы выбрали: KZT")
			return "KZT"
		case "TRY", "4":
			fmt.Println("Вы выбрали: TRY")
			return "TRY"
		default:
			fmt.Println("Такой валюты нет!")
			continue
		}
	}

}
func sumCurrency() float64 {
	input := ""
	sum := 0.0
	for {
		fmt.Println("Введите сумму для обмена:")
		fmt.Scan(&input)
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Введите числовое значение!")
			continue
		} else {
			sum = val
			break
		}

	}
	return sum
}
func currencyConverter() string {
	result := 0.0
	inputUserCurrency := fromCurrency()
	outoutUserCurrency := selectedCurrency()
	sum := sumCurrency()
	if inputUserCurrency == "USD" && outoutUserCurrency == "RUB" {
		result = sum * 83.01
	} else if inputUserCurrency == "USD" && outoutUserCurrency == "GBP" {
		result = sum * 0.78

	} else if inputUserCurrency == "USD" && outoutUserCurrency == "KZT" {
		result = sum * 447.14

	} else if inputUserCurrency == "USD" && outoutUserCurrency == "TRY" {

		result = sum * 34.20
	}
	if inputUserCurrency == "EUR" && outoutUserCurrency == "RUB" {
		result = sum * 90.11
	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "GBP" {
		result = sum * 0.85

	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "KZT" {
		result = sum * 485.20

	} else if inputUserCurrency == "EUR" && outoutUserCurrency == "TRY" {

		result = sum * 37.10
	}
	if inputUserCurrency == "CNY" && outoutUserCurrency == "RUB" {
		result = sum * 11.49
	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "GBP" {
		result = sum * 0.11

	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "KZT" {
		result = sum * 61.40

	} else if inputUserCurrency == "CNY" && outoutUserCurrency == "TRY" {

		result = sum * 4.70
	}
	if inputUserCurrency == "JPY" && outoutUserCurrency == "RUB" {
		result = sum * 0.53
	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "GBP" {
		result = sum * 0.0051

	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "KZT" {
		result = sum * 2.85

	} else if inputUserCurrency == "JPY" && outoutUserCurrency == "TRY" {

		result = sum * 0.22
	}

	resultF := fmt.Sprintf("Вы выбрали обменять: %s на %s\nИтоговая сумма которую вы получаете: %.2f", inputUserCurrency, outoutUserCurrency, result)
	return resultF
}
