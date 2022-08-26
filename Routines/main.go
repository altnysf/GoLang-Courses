package main

import (
	"testmod/routinesexample"
	"time"
)

func main() {
	go routinesexample.Example()
	time.Sleep(100 * time.Millisecond)
}
