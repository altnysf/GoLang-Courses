package Examples

import "fmt"

type Products struct {
	ProductId   int
	ProductName string
	UnitPrice   float64
}

func Example2() {
	product := new(Products)

	product.ProductId = 1
	product.ProductName = "Laptop"
	product.UnitPrice = 3000.0

	fmt.Println(
		"Product ID : ", product.ProductId,
		"Product Name : ", product.ProductName,
		"Unit Price : ", product.UnitPrice)
}
