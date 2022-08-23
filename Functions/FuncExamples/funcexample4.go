package funcexamples

func Example4() {
	sayHello("Hello", " ", "Go", " ", "Developer", ".\n")
}

func sayHello(messages ...string) {
	for _, message := range messages {
		print(message)
	}
}
