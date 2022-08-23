package main

import "fmt"

func main() {
	firstSlice := []int{1, 3, 5, 7, 9}

	fmt.Println(firstSlice[1:4])
	/*OUTPUTS	[3 5 7] 	*/

	fmt.Println(firstSlice[1:])
	/*OUTPUTS	[3 5 7 9] 	*/
	fmt.Printf("--------------------------------------------------\n")

	firstSlice = append(firstSlice, 11)
	fmt.Println(firstSlice)
	/* OUTPUTS [1 3 5 7 9 11] */
	fmt.Printf("--------------------------------------------------\n")

	fmt.Println(len(firstSlice)) // Lenght of Slice
	fmt.Println(cap(firstSlice)) // Capacity of Slice
	fmt.Printf("--------------------------------------------------\n")

}
