# file
a file lib for init filepath and simple use

## usage
```shell script
package main

import (
	"fmt"
	"github.com/gohouse/golib/file"
)

func main()  {
	var filepath = "/tmp/xxx.log"

	fmt.Println(file.FileExists(filepath))
}
```