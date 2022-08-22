package main

import (
	"fmt"
	"time"
)

func main() {

	// DATE & TIME OPERATIONS

	// Create Our "Date & Time" Data With Date() method
	t := time.Date(2022, time.July, 22, 21, 30, 0, 0, time.UTC)

	fmt.Printf("\n Go Launch at %s\n", t) // String Type Output

	fmt.Printf("\n *************************************** \n")

	// Get Current Time Information With "time.Now()" method

	now := time.Now()

	fmt.Printf("\n The Time Is : %s\n", now)

	fmt.Printf("\n *************************************** \n")

	// Year, Month, Day and Weekday Informations

	fmt.Println("The Year Is :", t.Year())
	fmt.Println("The Month Is :", t.Month())
	fmt.Println("The Day Is :", t.Day())
	fmt.Println("The Weekday Is :", t.Weekday())

	fmt.Printf("\n *************************************** \n")

	// Add 1 Day

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v,%v,%v,%v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	// Add 1 Month

	nextMonth := t.AddDate(0, 1, 1)
	fmt.Printf("Next Month is %v,%v,%v,%v\n",
		nextMonth.Weekday(),
		nextMonth.Month(),
		nextMonth.Day(),
		nextMonth.Year())

	// Add 1 Year

	nextYear := t.AddDate(1, 1, 1)
	fmt.Printf("Next Year is %v,%v,%v,%v\n",
		nextYear.Weekday(),
		nextYear.Month(),
		nextYear.Day(),
		nextYear.Year())

}
