package main

import (
	"fmt"
	"goSession/toolkit/utils"
)

func calcData() int {
	totalValue := utils.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	total := calcData()
	fmt.Println(total)
}
