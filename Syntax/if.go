// // Uncomment this entire file

package main

import (
	"errors"
	"fmt"
)

func main() {

	var someVar = 9

	if someVar > 10 {
		fmt.Println(someVar)
	}
	// if, else if, else
	if someVar > 100 {
		fmt.Println("Gt 100")
	} else if someVar == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Lt 100")
	}
}

func someFunction() error {
	return errors.New("some error")
}

//func main() {
//	err := someFunction()
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	//if err := someFunction(); err != nil {
//	//	fmt.Println(err.Error())
//	//}
//}
