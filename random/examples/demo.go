package main

import (
	"fmt"
	"github.com/gohouse/golib/random"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//randomDemo()
	randPerm()
	//fisherYates()
}

func randPerm()  {

	//var arr2 = []int{0,1,2,3,4}
	var res = [][]int{}
	var wg sync.WaitGroup
	wg.Add(20)
	for i:=0;i<20;i++{
		go func(arg int) {
			fmt.Println(arg)
			res = append(res, rand.Perm(5))
			wg.Done()
		}(i)
	}
	wg.Wait()
	//printArrInt(arr2, res)
	//time.Sleep(100*time.Millisecond)
}
func printArrInt(arr []int, newArr [][]int)  {
	for _,item := range arr {
		for _,tmp := range newArr{
			if tmp[0] == item {
				fmt.Println(tmp)
			}
		}
	}
}


func randomDemo()  {
	fmt.Println(random.Rand())
	fmt.Println(random.Random(12))
	fmt.Println(random.Random(12,random.T_ALL))
	fmt.Println(random.RandomBetween(6, 11))
	fmt.Println(random.RandomBetween(6,11, random.T_ALL))
}

func fisherYates()  {
	var arr2 = []string{"a","b","c","d","e","f"}
	var res = [][]string{}
	for i:=0;i<20;i++{
		go func(arr []string) {
			tmp := random.FisherYates(arr)
			//fmt.Println(res)
			res = append(res, tmp)
		}(arr2)
	}
	time.Sleep(time.Millisecond * 100)
	printArr(arr2, res)
}

func printArr(arr []string, newArr [][]string)  {
	for _,item := range arr {
		for _,tmp := range newArr{
			if tmp[0] == item {
				fmt.Println(tmp)
			}
		}
	}
}
