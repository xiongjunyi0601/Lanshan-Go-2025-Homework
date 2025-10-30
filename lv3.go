package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	n := 9
	fmt.Printf("%d的阶乘是%d\n", n, factorial(n))
}
