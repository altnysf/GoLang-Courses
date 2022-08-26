package main

import (
	"testmod/routinesexample"
	"time"
)

func main() {
	routinesexample.Example()
	time.Sleep(100 * time.Millisecond)
}
