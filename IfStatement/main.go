package main

import "fmt"

func main() {

	num1 := 21
	num2 := 21

	if num1 > num2 {
		fmt.Println("Number 1 is Greater than Number 2")
	} else if num1 == num2 {
		fmt.Println("Number 1 is Equal to Number 2")
	} else {
		fmt.Println("Number 1 is Smaller than Number 2")
	}
}
