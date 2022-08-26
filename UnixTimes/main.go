package main

import (
	"fmt"
	"time"
)

// CALCULATE THE UNIX TIME

// What is Unix Time ?
/*Unix time is a system for describing a point in time.
It is the number of seconds that have elapsed since the Unix epoch,excluding leap seconds.
The Unix epoch is 00:00:00 UTC on 1 January 1970.*/

func main() {
	fmt.Printf("Unix Time of Now : %v\n", time.Now().Unix())
	time.Sleep(20 * time.Second)
	fmt.Println("20 Seconds Later")
	fmt.Printf("Unix Time of Now : %v\n", time.Now().Unix())
}
