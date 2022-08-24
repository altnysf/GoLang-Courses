package Examples

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func Example1() {
	x := Person{FirstName: "Go", LastName: "Lang", Age: 2011}

	fmt.Println(x.FirstName, x.LastName, x.Age)

	time.Sleep(2 * time.Second)

	fmt.Println("-------------------- ** --------------------")

}
