package main

import "fmt"

//var name string
//var namePointer *string
//
//func main() {
//	fmt.Println(name)
//	fmt.Println(namePointer)
//	fmt.Println(&name)
//}

//var name string = "yang"
//var namePointer *string = &name
//var nameValue = *namePointer
//
//func main() {
//	fmt.Println(name)
//	fmt.Println(namePointer)
//	fmt.Println(nameValue)
//}

//func changeName(n string) {
//	n = strings.ToUpper(n)
//}
//
//func changeName1(n *string) {
//	*n = strings.ToUpper(*n)
//}
//
//func main() {
//	name := "Yang"
//	//changeName(name)
//	changeName1(&name)
//	fmt.Println(name)
//}

// A User is a user type
type User struct {
	ID          int
	Name, Email string
}

var u = User{
	ID:    1,
	Name:  "Yang",
	Email: "yang@gmail.com",
}

func updateEmail(u *User) {
	u.Email = "fang@gmail.com"
	fmt.Println("in update email: ", u.Email)
}

func main() {
	fmt.Println("Pointers!")
	updateEmail(&u)
	fmt.Println("Updated User: ", u)
}
