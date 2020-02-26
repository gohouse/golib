# date
a date lib for golang

## install
```shell script
go get github.com/gohouse/golib/date
```

## example
```go
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
```
运行结果
```shell script
2019-12-01
2019-11-01
2019-01-01
-----------------------------
2016-02-29 23:59:59
2016-02-01
2016-01-01
-----------------------------
2011-02-28 00:00:00
2011-02-01
2011-01-01
```