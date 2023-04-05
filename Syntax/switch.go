package main

import "fmt"

func main() {

	var city string

	switch city {
	case "Xian":
		fmt.Println("I live in Xian")
	case "Chengdu":
		fmt.Println("I live in Chengdu")
	case "Shanghai":
		fmt.Println("I live in Shanghai")
	default:
		fmt.Println("I'm not live in earth.")
	}

	var i = 9

	switch {
	case i != 10:
		fmt.Println("NE 10")
		fallthrough
	case i < 10:
		fmt.Println("Lt 10")
	case i > 10:
		fmt.Println("Gt 10")
	default:
		fmt.Println("Equals 10")
	}
}
