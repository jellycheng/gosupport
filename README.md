# gosupport
go support, go functions

## 下载依赖
```
go get -u github.com/jellycheng/gosupport
```

## demo
```
vi main.go
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
)

func main() {
	total := gosupport.IntSum(10, 30, 51)
	fmt.Println(total) //91
}

```

