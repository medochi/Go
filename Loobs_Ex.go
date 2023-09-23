package main

import "fmt"

// if dosen't work ?!
//-------------------------------------------
//func Fizzbuzz1() {
//	for i := 1; i <= 50; i++ {
//		if i%5 == 0 && i%3 == 0 {
//			fmt.Printf("%v:FizzBuzz\n", i)
//		} else if i%5 == 0 {
//			fmt.Printf("%v:Buzz\n", i)
//		} else if i%3 == 0 {
//			fmt.Printf("%v:FizzBuzz\n", i)
//		} else {
//			continue
//		}
//	}
//}

func Fizzbuzz() {
	for i := 1; i <= 50; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Printf("%v:FizzBuzz\n", i)
		case i%3 == 0:
			fmt.Printf("%v:Fizz\n", i)
		case i%5 == 0:
			fmt.Printf("%v:Buzz\n", i)
		}
	}
}

func main() {
	Fizzbuzz1()
}
