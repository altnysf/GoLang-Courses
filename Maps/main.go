package main

import "fmt"

func main() {
	// Basic Example
	myMap := make(map[int]string)

	myMap[4] = "BMW"
	myMap[7] = "Mercedes"
	fmt.Println(myMap)
	fmt.Println("-----------------------------------------------------")

	states := make(map[string]string)
	states["IST"] = "ISTANBUL"
	states["ANK"] = "ANKARA"
	states["ANT"] = "ANTALYA"
	states["IZM"] = "IZMIR"
	fmt.Println(states)
	selectedState := states["ANK"]
	fmt.Println("Selected State :", selectedState)
	fmt.Println("-----------------------------------------------------")

	// DELETE A DATA FROM MAP
	delete(states, "IZM")
	fmt.Println(states)
	fmt.Println("-----------------------------------------------------")

	// ADD A DATA TO MAP
	states["ERZ"] = "ERZURUM"
	fmt.Println(states)
	fmt.Println("-----------------------------------------------------")

	// PRINT ALL ELEMENTS WITH FOR RANGE
	for i, v := range states {
		fmt.Printf("%v : %v\n", i, v)
	}

}
