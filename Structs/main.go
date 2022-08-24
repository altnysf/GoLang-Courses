package main

import (
	"fmt"
	"time"
)

//STRUCT
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Products struct {
	ProductId   int
	ProductName string
	UnitPrice   float64
}

func main() {
	x := Person{FirstName: "Go", LastName: "Lang", Age: 2011}

	fmt.Println(x.FirstName, x.LastName, x.Age)

	time.Sleep(2 * time.Second)

	fmt.Println("-------------------- ** --------------------")

	ProductDetails()
}

func ProductDetails() {
	product := new(Products)

	product.ProductId = 1
	product.ProductName = "Laptop"
	product.UnitPrice = 3000.

	fmt.Println(
		"Product ID : ", product.ProductId,
		"Product Name : ", product.ProductName,
		"Unit Price : ", product.UnitPrice)
}
