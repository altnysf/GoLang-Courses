package main

import "fmt"

func main() {
	// Basic Array Example

	myArray := [3]int{}
	myArray[0] = 32
	myArray[1] = 64
	myArray[2] = 52

	fmt.Println(myArray)

	fmt.Println(" -------------------------- ")

	secondArray := [...]string{"red", "green", "blue", "brown"} // [...] -> used for automatic sizing.
	fmt.Println(secondArray)
	fmt.Println("Current Value : ", secondArray[2]) // blue
	secondArray[2] = "black"                        // SET the value of 2 indexes element
	fmt.Println("New Value : ", secondArray[2])     //black

	fmt.Println(len(secondArray)) // Lenght of Array
	fmt.Println(cap(secondArray)) // Capacity of Array
}
