package random

import "strings"

func Shuffle(arg string) (newStr string) {
	return strings.Join(FisherYates(strings.Split(arg, "")), "")
}
