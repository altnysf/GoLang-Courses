package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Value : ", i)
	}

	/* OUTPUTS
	Value :  0
	Value :  1
	Value :  2
	Value :  3
	Value :  4
	*/

	sum := 0

	for sum <= 10 {
		sum += 2
		fmt.Println("Result : ", sum)
	}

	/* OUTPUTS
	Result :  2
	Result :  4
	Result :  6
	Result :  8
	Result :  10
	Result :  12
	*/
}
