package main

import "fmt"

func Classify(age int) {
	switch {
	case age == 0:
		fmt.Println("you are newborn")
	case age >= 1 && age <= 3:
		fmt.Println("you are toddler")
	case age < 13:
		fmt.Println("you are child")
	case age > 18:
		fmt.Println("you are teenager")
	default:
		fmt.Println("you are adult")
	}
}

func main() {
	Classify(10)
}
