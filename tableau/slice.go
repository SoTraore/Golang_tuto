package main

import (
	"fmt"
	"log"
	"sort"
)

func myfunct(name *string) {
	*name = "Toto"
}

func main() {
	var myTable []string

	myTable = append(myTable, "Hello")
	myTable = append(myTable, "world")
	myTable = append(myTable, "terrible")

	sort.Strings(myTable)
	fmt.Println(myTable)

	numbers := []int{1, 2, 3, 4, 5}
	log.Println(numbers)
	log.Println(numbers[1:4])

	names := []string{"Mama", "Tata", "Toto"}
	log.Println(names)

	name := "Tata"

	myfunct(&name)

	log.Println("Here is the new name", name)
}
