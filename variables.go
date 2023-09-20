package main

import "fmt"

func main() {
	var age = 15  // I choose the type `int`
	var age2 = 12 // the program choose the type `int`
	age3 := 20    // the program choose the type `int`

	fmt.Println("varable using var name type = ", age) // 15
	fmt.Println("varable using var name = ", age2)     // 12
	fmt.Println("varable using := ", age3)             // 20

}
