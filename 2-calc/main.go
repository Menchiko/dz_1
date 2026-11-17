package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	resultCalc()
}
func swithFunc() string {
	fmt.Println("Выберите операцию-AVG-SUM-MED")
	txtUser := " "
	reader := bufio.NewReader(os.Stdin)
	txtUser, _ = reader.ReadString('\n')
	txtUser = strings.TrimSpace(txtUser)

	switch txtUser {
	case "AVG":
		txtUser = "AVG"
	case "SUM":
		txtUser = "SUM"
	case "MED":
		txtUser = "MED"
	default:
		fmt.Println("ERROR")
	}
	return txtUser
}
func resultCalc() {
	res := swithFunc()
	if res == "AVG" {
		avg()
	} else if res == "SUM" {
		sum()
	} else if res == "MED" {
		med()
	}
}

func inputUserNumbers() []int {
	inputUser := []int{}
	inputBufio := []string{}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Input")
		os.Exit(1)
	} else {
		input = strings.TrimSpace(input)
		inputBufio = strings.Split(input, ",")
		for _, val := range inputBufio {
			val = strings.TrimSpace(val)
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Ввел букву")
				os.Exit(1)
			} else {
				inputUser = append(inputUser, num)
			}
		}
	}
	return inputUser
}
func sum() {
	slice := []int{}
	slice = append(slice, inputUserNumbers()...)
	result := 0
	for _, val := range slice {
		result += val
	}
	fmt.Println(result)
}
func avg() {
	slice := []int{}
	slice = append(slice, inputUserNumbers()...)
	result := 0
	res := 0.0
	for _, val := range slice {
		result += val
	}
	res = float64(result)
	res /= float64(len(slice))
	fmt.Println(res)
}
func med() {
	slice := []int{}
	slice = append(slice, inputUserNumbers()...)
	result := 0.0
	sort.Ints(slice)
	if len(slice)%2 == 0 {
		result = float64((slice[len(slice)/2-1])+(slice[len(slice)/2])) / 2.0
	} else {

		result = float64(float64(slice[len(slice)/2]))
	}
	fmt.Println(result)
}
