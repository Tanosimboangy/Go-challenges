package kata

import ( "strings" )

func Points(games []string) int {
  result := 0
  for _, game := range games {
    str := strings.Split(game, ":")
    x, y := str[0], str[1]
    switch {
      case x > y:
        result += 3
      case x == y:
        result += 1
    }
  }
  return result
}

// I chose this solution because it is well written
// The fact that it uses switch makes it more suitable for this case
// I didn't know that I could use the split() method as well