package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "medor"

	m1 := make(map[int]User)

	u1 := User{
		FirstName: "Souleymane",
		LastName:  "Traore",
		Age:       24,
	}

	m1[1] = u1

	fmt.Println(myMap["dog"])
	fmt.Println(u1.GetFirstName())
	fmt.Println(m1[1].FirstName)
}
