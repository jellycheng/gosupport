package curl

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHttpGet(t *testing.T) {
	con, err := HttpGet("http://devapi.xxx.com")
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

func TestHttpPost(t *testing.T) {
	paramData := "name=李四&age=19"
	con, err := HttpPost("http://devapi.xxx.com/test.php?a=1&b=hi", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

// go test -run="TestHttpPostJson"
func TestHttpPostJson(t *testing.T) {
	paramData := `{"name":"张三","age":28}`
	con, err := HttpPostJson("http://devapi.xxx.com/test.php?a=2&b=hi", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

func TestHttpPostForm(t *testing.T) {
	paramData := url.Values{"username": {"admin"}, "pwd": {"123456"}}
	con, err := HttpPostForm("http://devapi.xxx.com/test.php", paramData)
	if err != nil {
		fmt.Println("请求失败", err)
	} else {
		fmt.Println(con)
	}
}

// go test -run=TestBase64Decode
func TestBase64Decode(t *testing.T) {
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	s2, _ := Base64DecodeV2([]byte(s))
	fmt.Println(string(s2)) // {"alg":"HS256","typ":"JWT"}

}
