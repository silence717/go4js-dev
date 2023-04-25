package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("bad happened")
	}
	return nil
}

func openFile() error {
	f, err := os.Open("demo.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {

	//num := 9
	//err := isGreaterThanTen(num)
	//if err != nil {
	//	//fmt.Println(fmt.Errorf("%d is less than 10", num))
	//	//panic(err)
	//	log.Fatalln(err)
	//}

	err := openFile()

	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

}
