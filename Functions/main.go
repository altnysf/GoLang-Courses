package main

import (
	"fmt"
	"time"
)

func main() {
	getName()
}

func getName() {
	var name string
	fmt.Print("Write Your First Name : ")

	fmt.Scanln(&name)
	time.Sleep(1 * time.Second)
	fmt.Println("Welcome, ", name)

}
