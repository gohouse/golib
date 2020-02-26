package random

// Rand 默认随机生成6-32位的随机字符串(长度类型皆随机)
// 如果传入不同的参数,则分别对应不同的函数
func Rand(args ...interface{}) string {
	switch len(args) {
	case 1:
		return Random(args[0].(int))
	case 2:
		if r,ok := args[1].(RandType);ok {
			return Random(args[0].(int), r)
		}
		return RandomBetween(args[0].(int),args[1].(int))
	case 3:
		return RandomBetween(args[0].(int),args[1].(int),args[2].(RandType))
	default:
		var rt = getRandType()
		var length = RandBetween(6, 32)
		return randomReal(length, rt)
	}
}

// Random 随机生成指定长度的随机字符串(类型可选或随机)
func Random(length int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	return randomReal(length, rt)
}

// RandomBetween 随机生成指定长度区间的随机字符串(类型可选或随机)
func RandomBetween(min, max int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	var length = RandBetween(min, max)
	return randomReal(length, rt)
}

func randomReal(length int, fill RandType) string {
	str := fill.String()
	var res string
	var i = length
	for i > 0 {
		res += getCharactorFromStr(str)
		i--
	}
	return res
}

func RandCapital(length ...int) string {
	if len(length)>0 {
		return Random(length[0], T_CAPITAL)
	}
	return RandomBetween(6,32, T_CAPITAL)
}

func RandLowercase(length ...int) string {
	if len(length)>0 {
		return Random(length[0], T_LOWERCASE)
	}
	return RandomBetween(6,32, T_LOWERCASE)
}

func RandString(length ...int) string {
	if len(length)>0 {
		return Random(length[0], T_CAPITAL_LOWERCASE)
	}
	return Random(RandBetween(6,32), T_CAPITAL_LOWERCASE)
}

func RandNumberic(length ...int) string {
	if len(length)>0 {
		return Random(length[0], T_NUMBERIC)
	}
	return Random(RandBetween(6,32), T_NUMBERIC)
}

func RandAll(length ...int) string {
	if len(length)>0 {
		return Random(length[0], T_ALL)
	}
	return Random(RandBetween(6,32), T_ALL)
}
