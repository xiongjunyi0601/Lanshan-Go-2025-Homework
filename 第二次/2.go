package main

import "fmt"

func calculator(op string) func(a, b int) int {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	default:
		return nil
	}
}

func main() {
	add := calculator("+")
	fmt.Println(add(2, 3))

	multiply := calculator("*")
	fmt.Println(multiply(2, 3))
}
