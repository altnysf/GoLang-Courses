package main

import "FuncMod/funcexamples"

func main() {
	funcexamples.Example1() // OUTPUT : Welcome, "Username"
	println(" ------------------------------------------------------ ")
	funcexamples.Example2() // OUTPUT : Hello Go!
	println(" ------------------------------------------------------ ")
	funcexamples.Example3() // OUTPUT : Added :  3 terms for a total of 12
	println(" ------------------------------------------------------ ")
	funcexamples.Example4() // OUTPUT : Hello Go Developer.
	println(" ------------------------------------------------------ ")
	funcexamples.Defering()
	/* OUTPUT :
	Connection Open : false
	Connected to database !
	Defering Disconnect !
	Connection Open : true
	Doing Something
	Disconnected to database !
	Connection Open : false
	*/
	println(" ------------------------------------------------------ ")
}
