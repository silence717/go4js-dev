package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		fmt.Println(i)
		i += 1
	}

	//for i := 1; i <= 100; i++ {
	//	fmt.Println(i)
	//}

	//use range
	name := "yangfang"

	for index, value := range name {
		fmt.Println(index, string(value))
	}

}
