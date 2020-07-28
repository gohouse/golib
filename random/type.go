package random

import (
	"fmt"
	"math/rand"
)

// RandType ...
type RandType int

const (
	// 大写字母
	TypeCAPITAL RandType = iota + 1
	// 小写字母
	TypeLOWERCASE = TypeCAPITAL << 1
	// 数字
	TypeNUMBERIC = TypeCAPITAL << 2
)

const (
	StrNumberic  = `0123456789`
	StrLowercase = `abcdefghijklmnopqrstuvwxyz`
	StrCapital   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

func (rt RandType) String() string {
	switch rt {
	case TypeCAPITAL:
		return StrCapital
	case TypeLOWERCASE:
		return StrLowercase
	case TypeNUMBERIC:
		return StrNumberic
	case TypeCAPITAL|TypeLOWERCASE:
		return fmt.Sprint(StrCapital, StrLowercase)
	case TypeCAPITAL|TypeNUMBERIC:
		return fmt.Sprint(StrCapital, StrNumberic)
	case TypeLOWERCASE|TypeNUMBERIC:
		return fmt.Sprint(StrLowercase, StrNumberic)
	case TypeCAPITAL|TypeLOWERCASE|TypeNUMBERIC:
		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
	}
	return ""
}

func getRandType() RandType {
	var ct = []RandType{
		// 大写字母
		TypeCAPITAL,
		// 小写字母
		TypeLOWERCASE,
		// 数字
		TypeNUMBERIC,
		// 大小写
		TypeCAPITAL|TypeLOWERCASE,
		// 大写+数字
		TypeCAPITAL|TypeNUMBERIC,
		// 小写+数字
		TypeLOWERCASE|TypeLOWERCASE,
	}
	l := len(ct)
	i := rand.Intn(l)
	return ct[i]
}
