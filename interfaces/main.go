package main

import "fmt"

type User struct {
	id      int
	Name    string
	Age     int
	Address string
}

func (u User) GetID() int {
	return u.id
}

func (u User) GetUserInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Address: %s", u.Name, u.Age, u.Address)
}

type UserMethods interface {
	GetID() int
	GetUserInfo() string
}

func main() {
	u := User{
		id:      1,
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	fmt.Println(u.GetUserInfo())
}
