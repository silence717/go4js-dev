package main

import "fmt"

func main() {

	//var myArray [5]int
	//var mySlice = make([]int, 5, 10)
	//
	//myArray[0] = 1
	//mySlice[0] = 2
	//
	//fmt.Println(myArray)
	//fmt.Println(mySlice)
	//fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))

	//slice1 := []int{1, 2, 3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	originalSlice := []int{1, 2, 3}
	destination := make([]int, len(originalSlice))
	fmt.Println("Before:", originalSlice, destination)

	mysteryValue := copy(destination, originalSlice)

	fmt.Println("After:", originalSlice, destination, mysteryValue)

}
