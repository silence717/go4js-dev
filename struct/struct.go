package main

import "fmt"

//type User struct {
//	ID    int
//	Name  string
//	Email string
//}

type User struct {
	ID          int
	Name, Email string
}

type Group struct {
	role       string
	users      []User
	newestUser User
}

func describeUser(user User) string {
	desc := fmt.Sprintf("Name: %s, Email: %s, ID: %d", user.Name, user.Email, user.ID)
	return desc
}

func describeGroup(group Group) string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s", group.role, len(group.users), group.newestUser.Name)
	return desc
}

func main() {
	user := User{ID: 1, Name: "yang", Email: "yang@gmail.com"}
	user2 := User{ID: 2, Name: "fang", Email: "fang@gmail.com"}

	group := Group{role: "admin", users: []User{user, user2}, newestUser: user2}

	userDescription := describeUser(user)
	groupDescription := describeGroup(group)

	fmt.Println(userDescription)
	fmt.Println(groupDescription)
}
