package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("Connection Open : %v\n", isConnected)
	DatabaseProcessing()
	fmt.Printf("Connection Open : %v\n", isConnected)
}

func DatabaseProcessing() {
	Connect()
	defer Disconnect()

	fmt.Println("Defering Disconnect !")
	fmt.Printf("Connection Open : %v\n", isConnected)
	fmt.Println("Doing Something")
}

func Connect() {
	isConnected = true
	fmt.Println("Connected to database !")
}

func Disconnect() {
	isConnected = false
	fmt.Println("Disconnected to database !")
}
