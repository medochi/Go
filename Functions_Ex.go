package main

import "fmt"

func main() {
	gretting("medochi")
	//--------------------------------------------------------------------------
	fmt.Println(message())
	//--------------------------------------------------------------------------
	fmt.Printf("the sum of 10 + 20 + 30 is %v\n", sumOfThree(10, 20, 30))
	//--------------------------------------------------------------------------
	fmt.Println("the number is", returnNum())
	//--------------------------------------------------------------------------
	a, b := returnTwoNums()
	//--------------------------------------------------------------------------
	fmt.Println("any three numbers", sumOfThree(returnNum(), a, b))

}

func gretting(name string) { // function greet a name
	fmt.Println("hello mr :", name)
}

func message() string { // function that return a message

	return "medochi is aweasome"
}
func sumOfThree(first, secound, third int) int { // function that sum three numbers
	return first + secound + third
}
func returnNum() int { // function that return a number
	return 29
}
func returnTwoNums() (int, int) { // function that return two numbers
	return 25, 10
}
