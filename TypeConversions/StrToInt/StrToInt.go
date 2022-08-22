package strtoint

import (
	"fmt"
	"strconv"
)

func Convert() {

	// String to Int Convert

	a := "10"
	b := 4

	x, _ := strconv.Atoi(a)

	// "_" is "empty identifier"

	result := x + b
	fmt.Println(result)
}
