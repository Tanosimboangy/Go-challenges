func Between(a, b int) []int {
  numbers := make([]int, b - a + 1)
  for i, _ := range numbers {
    numbers[i] = a + i
  } 
  return numbers
}

// I chose this solution because it shows the right to append an item into an array, which is what I tried but did not succed