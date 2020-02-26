package main

import (
	"fmt"
	"github.com/gohouse/golib/date"
)

func main() {
	newDate := date.NewDate()
	fmt.Println(newDate.YesterdayDate())
	fmt.Println(newDate.LastMonthStartDate())
	fmt.Println(newDate.YearStartDate())

	fmt.Println("-----------------------------")

	newDate = date.NewDate(date.BindDate("2016-03-01"))
	fmt.Println(newDate.YesterdayEndDateTime())
	fmt.Println(newDate.LastMonthStartDate())
	fmt.Println(newDate.YearStartDate())

	fmt.Println("-----------------------------")

	newDate = date.NewDate(date.BindDateTime("2011-03-01 14:22:21"))
	fmt.Println(newDate.YesterdayStartDateTime())
	fmt.Println(newDate.LastMonthStartDate())
	fmt.Println(newDate.YearStartDate())
}
