package main

import "fmt"

func countElementFrequency(arr []int) map[int]int {
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}
	return frequency
}

func main() {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	result := countElementFrequency(numbers)
	fmt.Println("元素出现的次数：")
	for num, count := range result {
		fmt.Printf("数字 %d 出现了 %d 次\n", num, count)
	}
}
