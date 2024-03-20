func CountPositivesSumNegatives(numbers []int) []int {
  
  pos := 0
  neg := 0
  
  for _, num := range numbers {
    if num > 0 {
      pos++
    } else if num < 0 {
      neg += num
    }
  }
  
  return []int{pos, neg}
}