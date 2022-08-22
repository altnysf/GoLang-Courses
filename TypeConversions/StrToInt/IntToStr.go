package strtoint

import (
	"fmt"
	"strconv"
)

func Convert1() {

	// Int to String Convert

	a := "10"
	b := 4

	result := a + " --- " + strconv.Itoa(b)

	fmt.Println(result)
}
