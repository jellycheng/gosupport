# gosupport
```
go support, go functions
本库用途： go常用方法、算法等的封装，提供基于go原生语法打造的类库
本库代码无第三方库依赖，仅依赖go内置函数、go语法

```

## 下载依赖
```
go get -u github.com/jellycheng/gosupport
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gosupport
```

## 调用示例
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
	
	fmt.Println(gosupport.Md5V1("hello world")) //5eb63bbbe01eeed093cb22bb8f5acdc3
}

```

