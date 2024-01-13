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
    
    // 把整数值生成对应的压缩码
    sCode := utils.NewMintCompress().Compress(1680) 
    fmt.Println(sCode) //Yt
    
    // md5加密
    fmt.Println(gosupport.Md5("hello world")) //5eb63bbbe01eeed093cb22bb8f5acdc3

    // 时间戳转日期格式: 2021-05-24 16:59:04
    stime := gosupport.Time()
    fmt.Println(gosupport.Timestamp2DateTime(int(stime), 1))
    // 按格式输出时间： 2021年05月24日 17时03分05秒
    fmt.Println(gosupport.DateT("Y年m月d日 H时i分s秒", *gosupport.TimeNowPtr()))
    // 判断文件是否存在
    if b := gosupport.FileExists("./xxx.go");b==true {
        fmt.Println("文件存在")
    } else {
        fmt.Println("文件不存在")
    }
	
}

```

## 拉取git仓库代码
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

## 自定义路由
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"net/http"
)

func main() {

	myRoute := gosupport.NewMyRoute()
	myRoute.GET("^/$", func(w http.ResponseWriter, r *http.Request, placeholder map[string]string, numStr []string) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("首页：" + r.URL.Path))
	})
	myRoute.GET("^/callback/wx/agentid123$", func(w http.ResponseWriter, r *http.Request, placeholder map[string]string, numStr []string) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("当前匹配的地址为：" + r.URL.Path))
	})

	myRoute.GET("^/callback/alipay/<payid>$", func(w http.ResponseWriter, r *http.Request, placeholder map[string]string, numStr []string) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		tmp := fmt.Sprintf("当前匹配的地址为：%s <br/> placeholder:%+v <br/> numStr:%+v <br/> %+v",
			r.URL.Path, placeholder, numStr, r.URL.Query())
		_, _ = w.Write([]byte(tmp))
	})

	myRoute.POST("^/userinfo/edit$", func(w http.ResponseWriter, r *http.Request, placeholder map[string]string, numStr []string) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("修改用户信息，当前匹配的地址为：" + r.URL.Path))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		isMatch, hander, placeholder, numStr := myRoute.MatchRoute(r.Method, r.URL.Path)
		if isMatch {
			hander(w, r, placeholder, numStr)
		} else {
			w.WriteHeader(200)
			_, _ = w.Write([]byte("未匹配到地址：" + r.URL.Path))
		}

	})
	_ = http.ListenAndServe(":9999", nil)
}

```

