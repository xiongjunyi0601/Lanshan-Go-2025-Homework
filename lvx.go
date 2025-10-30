package main

import "fmt"

func average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func main() {
	var num, sum, count int

	fmt.Println("请输入整数(输入0结束):")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
		count++
	}

	avg := average(sum, count)

	fmt.Printf("平均成绩为%.2f，", avg)
	if avg >= 60 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格")
	}
}
