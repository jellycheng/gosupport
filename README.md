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

## 用户ID生成唯一邀请码
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport/utils"
)

func main() {
	inviteObj := utils.NewGetInviteCode()
	var userid int64 = 1234567
	code := inviteObj.EnCode(userid)
	fmt.Println(code) //deahx

}

```

## 日期时间
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"time"
)

func main() {
	// 今天：2024-04-25
	today := gosupport.Now().Format(gosupport.DateFormat)
	fmt.Println("今天：", today)

	// 明天：2024-04-26
	tomorrow := gosupport.Tomorrow().Format(gosupport.DateFormat)
	fmt.Println("明天：", tomorrow)

	// 昨天：2024-04-24
	yesterday := gosupport.Yesterday().Format(gosupport.DateFormat)
	fmt.Println("昨天：", yesterday)
	// 昨天时间戳：1713924837
	yesterdayTimestamp := gosupport.Yesterday().Timestamp()
	fmt.Println("昨天时间戳:", yesterdayTimestamp)
	// 本周一日期：2024-04-22
	fmt.Println("本周一日期：", gosupport.Now().ToMondayString())
	// 本周日日期：2024-04-29
	fmt.Println("本周日日期：", gosupport.Now().ToSundayString())
	// 上周周一日期：2024-04-15
	fmt.Println("上周周一日期:", gosupport.PrevWeekMonday())
	// 上周周日日期：2024-04-21
	fmt.Println("上周周日日期:", gosupport.PrevWeekSunday())

	// 本月1号: 2024-04-01
	now := time.Now()
	month1 := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, gosupport.GetShanghaiTimezone()).Format(gosupport.DateFormat)
	fmt.Println("本月1号:", month1)

	// 本月最后1天日期：2024-04-30
	now2 := time.Now()
	month2 := time.Date(now2.Year(), now2.Month(), gosupport.GetMonthDay(now.Year(), int(now.Month())), 0, 0, 0, 0, gosupport.GetShanghaiTimezone()).Format(gosupport.DateFormat)
	fmt.Println("本月最后1天日期:", month2)

	// 获取当前季度: 2
	fmt.Println("获取当前季度: ", gosupport.Now().Quarter())

}


```

## web上传示例
```
package main

import (
	"fmt"
	"github.com/jellycheng/gosupport/curl"
	"net/http"
	"strconv"
)

// w响应请求对象，r接收请求对象
func UploadFileFunc(w http.ResponseWriter, r *http.Request) {
	upObj := curl.NewUploadFile()
	// 接收上传文件
	f, err := upObj.FormFile(r, "upfile")
	if err == nil {
		fmt.Println("文件名：" + f.Filename)
		fmt.Println("文件大小：" + strconv.FormatInt(f.Size, 10))
		// 保存文件
		_ = upObj.SaveUploadedFile(f, "./"+f.Filename)
	} else {
		fmt.Println("err:" + err.Error())
	}

	_, _ = w.Write([]byte("响应内容：hi，你好"))
}


func main() {
	listenIpPort := ":8080" // 监听端口
	// 配置路由
	http.HandleFunc("/", UploadFileFunc)
	err := http.ListenAndServe(listenIpPort, nil)
	if err != nil {
		fmt.Printf("服务启动失败, err:%v\n", err)
	} else {
		fmt.Println("服务启动成功, server:" + listenIpPort)
	}
}


```

