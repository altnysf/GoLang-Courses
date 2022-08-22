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
}
