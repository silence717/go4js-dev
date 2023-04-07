package main

import "fmt"

func main() {
	var userEmails = make(map[int]string)

	userEmails[1] = "yf@gmail.com"
	userEmails[2] = "yf2023@gmail.com"

	fmt.Println(userEmails)

	//userEmails := map[int]string{
	//	1: "yf@gmail.com",
	//	2: "yf2023@gmail.com",
	//}
	//
	//userEmails[1] = "yf3@gmail.com"
	//
	//fmt.Println(userEmails)
	//
	//firstEmail, ok := userEmails[1]
	////firstEmail, ok := userEmails[4]
	//fmt.Println(firstEmail, ok)
	//
	//if email, ok := userEmails[4]; ok {
	//	fmt.Println(email)
	//} else {
	//	fmt.Println("email not existed")
	//}
	//
	//delete(userEmails, 2)
	//
	//fmt.Println(userEmails)

	for k, v := range userEmails {
		fmt.Printf("%s key is %d.\n", v, k)
	}
}
