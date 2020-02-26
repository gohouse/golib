package random

import (
	"math/rand"
)

func getRandType() RandType {
	var ct = []RandType{
		// 大写字母
		T_CAPITAL,
		// 小写字母
		T_LOWERCASE,
		// 数字
		T_NUMBERIC,
		// 小写字母+数字
		T_LOWERCASE_NUMBERIC,
		// 大写字母+数字
		T_CAPITAL_NUMBERIC,
		// 大写字母+小写字母
		T_CAPITAL_LOWERCASE,
		// 数字+字母
		T_ALL,
	}
	l := len(ct)
	i := rand.Intn(l)
	return ct[i]
}

func getCharactorFromStr(str string) string {
	strLen := len(str)
	return string(([]rune(str))[rand.Intn(strLen-1)])
}

//func getFillStr(rt RandType) string {
//	switch rt {
//	case T_CAPITAL:
//		return StrCapital
//	case T_LOWERCASE:
//		return StrLowercase
//	case T_NUMBERIC:
//		return StrNumberic
//	case T_CAPITAL_LOWERCASE:
//		return fmt.Sprint(StrCapital, StrLowercase)
//	case T_CAPITAL_NUMBERIC:
//		return fmt.Sprint(StrCapital, StrNumberic)
//	case T_LOWERCASE_NUMBERIC:
//		return fmt.Sprint(StrLowercase, StrNumberic)
//	case T_ALL:
//		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
//	}
//	return ""
//}

func RandBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}
