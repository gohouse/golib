# varBindValue

给变量指针绑定值
```go

package main

import (
	"github.com/gohouse/gocar/varBindValue"
	"fmt"
)
func main() {
	var a int

	err := varBindValue.BindVal(&a, 234)

	fmt.Println(a)
	fmt.Println(err)
}
```
结果
```bash
234
<nil>
```
