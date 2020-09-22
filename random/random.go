package random

import (
	"fmt"
	"math/rand"
)

type iRandom interface {
	Rand(length int, fill RandType) string
	RandString(length int) string
	RandBetween(min, max int) int
	RandVariable(length int) string
}

// RandType ...
type RandType int

const (
	// 大写字母
	TypeCAPITAL RandType = 1
	// 小写字母
	TypeLOWERCASE = TypeCAPITAL << 1
	// 数字
	TypeNUMBERIC = TypeCAPITAL << 2
	// 所有
	TypeALL = TypeCAPITAL | TypeLOWERCASE | TypeNUMBERIC
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
	case TypeCAPITAL | TypeLOWERCASE:
		return fmt.Sprint(StrCapital, StrLowercase)
	case TypeCAPITAL | TypeNUMBERIC:
		return fmt.Sprint(StrCapital, StrNumberic)
	case TypeLOWERCASE | TypeNUMBERIC:
		return fmt.Sprint(StrLowercase, StrNumberic)
	case TypeCAPITAL | TypeLOWERCASE | TypeNUMBERIC:
		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
	}
	return ""
}

// GetRandTypeList 获取给定的字符串类型列表
func GetRandTypeList() []RandType {
	return []RandType{
		// 大写字母
		TypeCAPITAL,
		// 小写字母
		TypeLOWERCASE,
		// 数字
		TypeNUMBERIC,
	}
}

// Rand 根据给定字符串类型获取指定长度的随机字符串
func Rand(length int, fill RandType) (res string) {
	if length == 0 {
		return ""
	}
	for length > 0 {
		length--
		res += getCharactorFromStr(fill.String())
	}
	return
}

// RandString 随机长度字符串
func RandString(length int) string {
	return Rand(length, TypeALL)
}

// RandBetween [min,max]
func RandBetween(min, max int) int {
	if max < min {
		return 0
	}
	return rand.Intn(max-min+1) + min
}

// RandVariable 随机变量名字,不能以数字开头
func RandVariable(length int) string {
	if length == 0 {
		return ""
	}
	return fmt.Sprint(Rand(1, TypeLOWERCASE|TypeCAPITAL), Rand(length-1, 7))
}

func getCharactorFromStr(str string) string {
	strLen := len(str)
	return string(([]rune(str))[rand.Intn(strLen-1)])
}
