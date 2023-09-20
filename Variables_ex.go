package main

import "fmt"

func main() {
	var favColour = "yellow"
	birthYear, age := 2005, 18
	var (
		firstName  = "Mohammed "
		secoudName = "gamil"
	)
	var ageDays int
	ageDays = 365 * age
	fmt.Println("hello my name is", firstName+secoudName)
	fmt.Printf("i was porn in %v and  my age is %v\n", birthYear, age)
	fmt.Println("my age in days is ", ageDays)
	fmt.Println("my favourite color is", favColour)
}
