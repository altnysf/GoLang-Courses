package funcexamples

func Example3() {
	numTerms, sum := add(3, 4, 5)
	println("Added : ", numTerms, "terms for a total of", sum)
}

// Variadic

func add(terms ...int) (int, int) {
	result := 0

	for _, term := range terms {
		result += term
	}
	return len(terms), result
}
