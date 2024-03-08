package main

import "fmt"

func Grow(arr []int) int{
  sum := 1
  for _, number := range arr {
    sum = sum * number
  }
  return sum
}

func main() {
	result := Grow([1, 2, 3, 4])
	fmt.Println(result)
}