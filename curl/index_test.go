package curl

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHttpGet(t *testing.T) {
	con,err := HttpGet("http://devapi.nfangbian.com")
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

func TestHttpPost(t *testing.T) {
	paramData := "name=李四&age=19"
	con,err := HttpPost("http://devapi.nfangbian.com/test.php?a=1&b=hi", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

// go test -run="TestHttpPostJson"
func TestHttpPostJson(t *testing.T) {
	paramData := `{"name":"张三","age":28}`
	con,err := HttpPostJson("http://devapi.nfangbian.com/test.php?a=2&b=hi", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

func TestHttpPostForm(t *testing.T) {
	paramData := url.Values{"username": {"admin"}, "pwd":{"123456"}}
	con,err := HttpPostForm("http://devapi.nfangbian.com/test.php", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

// go test -run="TestParseUrlAndAppendParam"
func TestParseUrlAndAppendParam(t *testing.T) {
	s,_ := ParseUrlAndAppendParam("http://www.nfangbian.com?a=yes&x=333&username=admin#wechat", map[string][]string{
		"a":[]string{"hello", "good"},
		"a3":[]string{"xxxx"},
	})
	// http://www.nfangbian.com?a=yes&a=hello&a=good&a3=xxxx&username=admin&x=333#wechat
	fmt.Println(s)
}
