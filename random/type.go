package random

import "fmt"

// RandType ...
type RandType int

const (
	// 大写字母
	T_CAPITAL RandType = iota + 1
	// 小写字母
	T_LOWERCASE
	// 数字
	T_NUMBERIC
	// 小写字母+数字
	T_LOWERCASE_NUMBERIC
	// 大写字母+数字
	T_CAPITAL_NUMBERIC
	// 大写字母+小写字母
	T_CAPITAL_LOWERCASE
	// 数字+字母
	T_ALL
)

const (
	StrNumberic  = `0123456789`
	StrLowercase = `abcdefghijklmnopqrstuvwxyz`
	StrCapital   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

func (rt RandType) String() string {
	switch rt {
	case T_CAPITAL:
		return StrCapital
	case T_LOWERCASE:
		return StrLowercase
	case T_NUMBERIC:
		return StrNumberic
	case T_CAPITAL_LOWERCASE:
		return fmt.Sprint(StrCapital, StrLowercase)
	case T_CAPITAL_NUMBERIC:
		return fmt.Sprint(StrCapital, StrNumberic)
	case T_LOWERCASE_NUMBERIC:
		return fmt.Sprint(StrLowercase, StrNumberic)
	case T_ALL:
		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
	}
	return ""
}
