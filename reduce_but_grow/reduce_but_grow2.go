package main

func Grow2(arr []int) int {

	var result = 1

	for _, item := range arr {

		if item == 0 {
			return 0
		}

		result *= item
	}

	return result
}
