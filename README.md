# gosupport
```
go support, go functions
本库用途： go常用方法、算法等的封装，提供基于go原生语法打造的类库
本库代码无第三方库依赖，仅依赖go内置函数、go语法

常用算法、工具是日积月累的，持续打磨，完善，坚持、加油！！！

```
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Reference](https://pkg.go.dev/badge/github.com/jellycheng/gosupport.svg)](https://pkg.go.dev/github.com/jellycheng/gosupport)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/jellycheng/gosupport/badges/download-count.svg)](https://goproxy.cn/stats/github.com/jellycheng/gosupport/badges/download-count.svg)

## Requirements
```
gosupport library requires Go version >=1.14

```

## Install下载依赖
```
go get -u github.com/jellycheng/gosupport
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gosupport

```

## Documentation
[https://pkg.go.dev/github.com/jellycheng/gosupport](https://pkg.go.dev/github.com/jellycheng/gosupport)

## Usage调用示例
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
    "github.com/jellycheng/gosupport/utils"
)

func main() {
    // 求和
	total := gosupport.IntSum(10, 30, 51)
	fmt.Println(total) //91
    
    sCode := utils.NewMintCompress().Compress(1680) // 数值生成对应的压缩码
	fmt.Println(sCode) //Yt
	// md5加密
	fmt.Println(gosupport.Md5V1("hello world")) //5eb63bbbe01eeed093cb22bb8f5acdc3

    // 时间戳转日期格式: 2021-05-24 16:59:04
    stime := gosupport.Time()
    fmt.Println(gosupport.Timestamp2DateTime(int(stime), 1))
    // 2021年05月24日 17时03分05秒
    fmt.Println(gosupport.DateT("Y年m月d日 H时i分s秒", *gosupport.TimeNowPtr()))

}

```

