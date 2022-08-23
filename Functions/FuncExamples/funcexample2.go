package funcexamples

import "fmt"

func Example2() {
	a, b := swap("Go!", "Hello")

	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
