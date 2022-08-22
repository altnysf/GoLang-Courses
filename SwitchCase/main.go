package main

import "fmt"

func main() {

	var score float64

	fmt.Print("Enter your score for your final exam : ")
	fmt.Scanf("%v", &score)

	switch {
	case score >= 0 && score <= 45:
		fmt.Println("Your Grade is -F-")

	case score > 45 && score <= 59:
		fmt.Println("Your Grade is -E-")

	case score > 59 && score <= 69:
		fmt.Println("Your Grade is -D-")

	case score > 69 && score <= 79:
		fmt.Println("Your Grade is -C-")

	case score > 79 && score <= 89:
		fmt.Println("Your Grade is -B-")

	case score > 89 && score <= 100:
		fmt.Println("Your Grade is -A-")

	default:
		fmt.Println("Please Enter a Score Between 0 and 100")
	}
}
