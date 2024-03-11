package main

import (
  "strings"
)

func main(str string) string {
  words := strings.Split(str, " ")
  var reversedWords []string
  
  for _, item := range words {
    if item == "" {
      reversedWords = append(reversedWords, item)
    } else {
      reverseSingleString := reverseWord(item)
      reversedWords = append(reversedWords, reverseSingleString)
    }
  }
  return strings.Join(reversedWords, " ")
}

func reverseWord(str string) string {
  tmp := []byte{}
  for i := len(str) - 1; i >= 0; i-- {
    tmp = append(tmp, str[i])
  }
  return string(tmp)
}