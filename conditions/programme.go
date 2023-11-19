package main

import (
	"fmt"
	"log"
)

func main() {

	var myBool bool

	myBool = false
	isko := true

	if myBool {
		log.Println(myBool)
	} else if isko {
		log.Println(isko)
	} else {
		log.Println("else")
	}

	var input int
	input = 100

	switch input {
	case 1:
		log.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		log.Println("C")
	case 100:
		log.Println("Right")
	default:
		fmt.Println("Default")
	}

}
