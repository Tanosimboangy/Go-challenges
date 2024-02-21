package kata

import (
	"strconv"
)

func StringToNumber(str string) int {
	item, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return item
}
