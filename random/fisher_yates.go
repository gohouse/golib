package random

import (
	"math/rand"
)

// FisherYates 洗牌算法,确保已经 seed 种子
func FisherYates(arr []interface{}) []interface{} {
	var newArr []interface{}
	for _, item := range arr {
		newArr = append(newArr, item)
	}
	//rand.Seed(time.Now().UnixNano())
	for i := len(newArr) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		newArr[i], newArr[num] = newArr[num], newArr[i]
	}
	return newArr
}
