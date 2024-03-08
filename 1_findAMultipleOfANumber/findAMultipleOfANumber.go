package main

import "fmt"

func findMultiples(integer, limit int) []int {
  var nil []int
  for i := integer; i <= limit; i++ {
    if i%integer == 0 {
      nil = append(nil, i)
    }
  }
  return nil
}

func main() {
	result := findMultiples(2, 20)
	fmt.Println(result)
}