package main

import "fmt"

func main() {

	// Variable Definition Method 1

	var exampleVariable string
	exampleVariable = "Hello Go! - Variable 1"
	fmt.Println(exampleVariable)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Variable Definition Method 2

	var exampleText = "Hello Go! - Variable 2"
	fmt.Println(exampleText)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Variable Definition Method 3
	var a, b, c int
	a = 4
	b = 5
	c = 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Variable Definition Method 4
	var d, e, f int = 15, 35, 56

	fmt.Println(d, e, f)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Variable Definition Method 5
	var j, k, l, m = 15, true, 5.6, "Example String"

	fmt.Println(j, k, l, m)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Variable Definition Method 6
	x, y, z := true, 5.6, "String Variable"

	fmt.Println(x, y, z)

	fmt.Println("---------------------------------------------------------------------------- ")
}
