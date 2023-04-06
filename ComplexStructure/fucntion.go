package main

import "fmt"

//func printName(name string) string {
//	fmt.Println(name)
//	return name
//}
//
//func main() {
//	printName("yang")
//}

// multiple returns
//func myFunc() (int, int) {
//	return 1, 2
//}
//
//func main() {
//	x, y := myFunc()
//	fmt.Println(x, y)
//}

// named return
//func myFunc() (ageOfCat int, ageOfDog int) {
//	ageOfCat = 3
//	ageOfDog = 5
//	return
//}
//
//func main() {
//	x, y := myFunc()
//	fmt.Println(x, y)
//}

//func average(num1, num2, num3 float64) float64 {
//	total := num1 + num2 + num3
//	return total / 3
//}

//func average(numbers ...float64) float64 {
//	total := 0.0
//	for _, number := range numbers {
//		total += number
//	}
//	return total / float64(len(numbers))
//}
//
//func main() {
//	fmt.Println(average(10, 5, 7))
//}

// Functions As Types
//func isOlderThanTen(age int) bool {
//	return age > 10
//}
//
//func isNotZero(num int) bool {
//	return num != 0
//}
//
//var myFuncVariable func(int) bool
//
//func main() {
//	myFuncVariable = isOlderThanTen
//	fmt.Println(myFuncVariable(9))
//
//	myFuncVariable = isNotZero
//	fmt.Println(myFuncVariable(9))
//}

// IIFE(Immediately Invoked Function Expressions) like js
var age = 9

var isOlderThanTen = func() bool {
	return age > 10
}()

func main() {
	fmt.Println(isOlderThanTen)
}
