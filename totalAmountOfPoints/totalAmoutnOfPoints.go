package kata

func Points(games []string) int {
  var points = 0
	for _, v := range games {
		if v[0] > v[2] {
        points += 3
    } else if v[0] == v[2] {
        points += 1
    }
	}
    return points
}