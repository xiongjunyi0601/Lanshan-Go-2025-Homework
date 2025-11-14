package main

import "fmt"

type BasketballPlayer struct {
	Name     string
	Number   int
	Position string
}

func (p BasketballPlayer) Equal(other BasketballPlayer) bool {
	return p.Name == other.Name && p.Number == other.Number
}

func main() {
	players := []BasketballPlayer{}

	players = append(players, BasketballPlayer{
		Name:     "Stephen Curry",
		Number:   30,
		Position: "PG",
	})
	players = append(players, BasketballPlayer{
		Name:     "Kyrie Irving",
		Number:   11,
		Position: "PG",
	})
	fmt.Println("已添加球员：Stephen Curry、Kyrie Irving")

	target := BasketballPlayer{Name: "Stephen Curry", Number: 30}
	fmt.Printf("\n查询球员【%s（%d号）】：\n", target.Name, target.Number)
	for _, p := range players {
		if p.Equal(target) {
			fmt.Printf("姓名：%s | 号码：%d | 位置：%s\n",
				p.Name, p.Number, p.Position)
		}
	}
}
