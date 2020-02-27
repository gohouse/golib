# structAutoGenerate

根据struct, 自动生成对应字段的`Set()`和`Get()`方法, 同时会为这些方法生成对应的接口

```go
package main

import (
	"github.com/gohouse/gocar/structAutoGenerate"
	"fmt"
)

type Users struct {
    Uid  int    `gorose:"uid"`
    Name string `gorose:"name"`
    Age  int    `gorose:"age"`
}
func main() {
    //var a tt
    err := structAutoGenerate.New(&structAutoGenerate.Option{
        // 要生成的机构体对象
        Obj:         Users{},
        // 生成文件包名
        PackageName: "user",
        // 指定生成文件保存目录
        SavePath: "./users.go",
    }).Generate()
    
    fmt.Println(err)
}
```
生成
```go
package user

type IUsers interface {
	SetUid(arg int)
	GetUid() int
	SetName(arg string)
	GetName() string
	SetAge(arg int)
	GetAge() int
}
type Users struct {
	Uid  int    `gorose:"uid"`
	Name string `gorose:"name"`
	Age  int    `gorose:"age"`
}

func NewUsers() *Users {
	return new(Users)
}

func (o *Users) SetUid(arg int) {
	o.Uid = arg
}

func (o *Users) GetUid() int {
	return o.Uid
}

func (o *Users) SetName(arg string) {
	o.Name = arg
}

func (o *Users) GetName() string {
	return o.Name
}

func (o *Users) SetAge(arg int) {
	o.Age = arg
}

func (o *Users) GetAge() int {
	return o.Age
}
```