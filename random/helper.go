package random

import (
	"math/rand"
)

func getCharactorFromStr(str string) string {
	strLen := len(str)
	return string(([]rune(str))[rand.Intn(strLen-1)])
}

func RandBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}
