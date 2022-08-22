package main

import "fmt"

func main() {

	// Değişken Tanımlama Yöntemi 1

	var degisken string
	degisken = "Merhaba Go! - Değişken 1"
	fmt.Println(degisken)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Değişken Tanımlama Yöntemi 2

	var degisken1 = "Merhaba Go! - Değişken 2"
	fmt.Println(degisken1)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Değişken Tanımlama Yöntemi 3
	var a, b, c int
	a = 4
	b = 5
	c = 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Değişken Tanımlama Yöntemi 4
	var d, e, f int = 15, 35, 56

	fmt.Println(d, e, f)

	fmt.Println("---------------------------------------------------------------------------- ")

	// Değişken Tanımlama Yöntemi 5
	var j, k, l, m = 15, true, 5.6, "Metin"

	fmt.Println(j, k, l, m)

	fmt.Println("---------------------------------------------------------------------------- ")
}
