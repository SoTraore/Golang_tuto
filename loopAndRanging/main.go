package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i + 1)
	}

	animals := []string{"cat", "fish", "dog", "dragon"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	myString := "This is my personal string"

	for i, l := range myString {
		log.Println(i, l)
	}

	var users []User

	users = append(users, User{"Souleymane", "TRAORE", "souleymane@etu.fr", 24})
	users = append(users, User{"Yacouni", "TOURE", "yacouni@etu.fr", 25})
	users = append(users, User{"Daoulata", "TRAORE", "daoulata@etu.fr", 20})

	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

	myMap := make(map[int]string)

	myMap[1] = "One"
	myMap[2] = "Two"
	myMap[3] = "Three"

	for key, value := range myMap {
		fmt.Println(key, value)
	}

}
