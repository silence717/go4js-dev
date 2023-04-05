package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 4
	var toggle = true
	name := "yang"
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(toggle))
	//fmt.Println(reflect.TypeOf(float64(x) * 5.5))
}
