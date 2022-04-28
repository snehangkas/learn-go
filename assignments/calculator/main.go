package main

import (
	"fmt"
)

func main() {
	var result float64
	result = doCalculate(result, 100, "+")
	fmt.Println(result)
	result = doCalculate(result, 50, "-")
	fmt.Println(result)
	result = doCalculate(result, 20, "/")
	fmt.Println(result)
	result = doCalculate(result, 10, "*")
	fmt.Println(result)
}

func doCalculate(result, operand float64, operator string) float64 {
	switch operator {
	case "+":
		return result + operand
	case "-":
		return result - operand
	case "*":
		return result * operand
	case "/":
		return result / operand
	default:
		return result // ideally should
	}
}
