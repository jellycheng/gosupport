# gosupport
```
go support, go functions
本库用途： go常用方法、算法等封装，提供原生工具
本库代码无第三方库依赖，仅依赖go内置语法

```

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

