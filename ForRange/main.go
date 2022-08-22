package main

import "fmt"

func main() {

	products := []string{"Laptop", "Desktop", "Mouse", "Keyboard", "Monitor"} // It's a "Slice"

	for i, v := range products {
		fmt.Println("Index is ", i, " and Product Name is ", v)
	}

	/*OUTPUTS
	Index is  0  and Product Name is  Laptop
	Index is  1  and Product Name is  Desktop
	Index is  2  and Product Name is  Mouse
	Index is  3  and Product Name is  Keyboard
	Index is  4  and Product Name is  Monitor
	*/

	numbers := [...]int{1, 2, 3, 4, 5}

	for i := range numbers {
		fmt.Println("Array Item ", i, " Value is ", numbers[i])
	}

	/*OUTPUTS
	Array Item  0  Value is  1
	Array Item  1  Value is  2
	Array Item  2  Value is  3
	Array Item  3  Value is  4
	Array Item  4  Value is  5
	*/

}
