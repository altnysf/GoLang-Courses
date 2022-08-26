package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, time.August, 8, 26, 25, 0, 0, time.UTC)
	fmt.Printf("\nWorking : %s\n", t)

	fmt.Println("----------------------------------------------")

	now := time.Now()
	fmt.Printf("Current Time : %s\n", now)

	fmt.Println("----------------------------------------------")

	fmt.Println("Month :", now.Month())
	fmt.Println("Day :", now.Day())
	fmt.Println("Day of Week :", now.Weekday())
	fmt.Println("Year :", now.Year())

	fmt.Println("----------------------------------------------")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v\n", tomorrow.Year(), tomorrow.Month(), tomorrow.Weekday(), tomorrow.Day())

	fmt.Println("----------------------------------------------")

	longFormat := "Friday, August 27, 2022"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

	fmt.Println("----------------------------------------------")

	shortFormat := "08/27/2022"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortFormat))
}
