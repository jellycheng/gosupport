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

直接获取master分支代码：
    go get -u github.com/jellycheng/gosupport@master
    
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
    
    sCode := utils.NewMintCompress().Compress(1680) // 把整数值生成对应的压缩码
    fmt.Println(sCode) //Yt
    // md5加密
    fmt.Println(gosupport.Md5V1("hello world")) //5eb63bbbe01eeed093cb22bb8f5acdc3

    // 时间戳转日期格式: 2021-05-24 16:59:04
    stime := gosupport.Time()
    fmt.Println(gosupport.Timestamp2DateTime(int(stime), 1))
    // 2021年05月24日 17时03分05秒
    fmt.Println(gosupport.DateT("Y年m月d日 H时i分s秒", *gosupport.TimeNowPtr()))
    // 判断文件是否存在
    b := gosupport.FileExists("./main6.go")
    if b {
        fmt.Println("文件存在")
    } else {
        fmt.Println("文件不存在")
    }
	
}

```

## 拉取git仓库代码示例
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
)

func main() {
	// err := gosupport.GitCloneIntoPath("git仓库地址", "分支名", "代码保存目录")
	err := gosupport.GitCloneIntoPath("https://github.com/jellycheng/cjsRedis.git", "master", "./cjs")
	if err == nil {
		fmt.Println("代码拉取完成")
	} else {
		fmt.Println("代码拉取失败：", err.Error())
	}
}

```
