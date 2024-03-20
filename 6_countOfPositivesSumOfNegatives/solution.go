func CountPositivesSumNegatives(numbers []int) []int {

	var pos, neg int
	for _, value := range numbers {
		switch {
		case value > 0:
			pos++
		case value < 0:
			neg += value
		}
	}
	return []int{pos, neg}
}

// I chose this solution as it uses switch instead of if statement
// It's kind of fancy to see switch being used here