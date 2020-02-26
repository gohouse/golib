package main

import (
	"fmt"
	"github.com/gohouse/golib/file"
)

func main()  {
	var filepath = "/tmp/xxx.log"

	f := file.NewFile(filepath)
	fmt.Println(f.Exists())
}
