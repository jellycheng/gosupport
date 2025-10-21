package curl

import (
	"fmt"
	"testing"
)

func TestNewParams(t *testing.T) {

	pObj := NewParams()

	//添加参数
	pObj.Set("userid", 123)
	pObj.Set("username", "tom")
	pObj.Set("is_supper", true)
	pObj.Set("price", 100.25)
	pObj.Set("empty", nil)

	println(pObj.ToQueryString())

	for k, v := range pObj {
		println(k, ParamValueToString(v))
	}

}

// go test -run="TestAppendParamToUrl"
func TestAppendParamToUrl(t *testing.T) {
	fmt.Println(AppendParamToUrl("http://xxx.com:80/fangan/index?xyz=88#hello", "a=你好&b=goods"))
}

// go test -run=TestParseQueryParams
func TestParseQueryParams(t *testing.T) {
	fmt.Println(ParseQueryParams("app_id=1128&userid=456").Get("app_id"))
}

// go test -run="TestParseUrlAndAppendParam"
func TestParseUrlAndAppendParam(t *testing.T) {
	s, _ := ParseUrlAndAppendParam("http://www.xxx.com?a=yes&x=333&username=admin#wechat",
		map[string][]string{
			"a":  []string{"hello", "good"},
			"a3": []string{"xxxx"},
		})
	// http://www.xxx.com?a=yes&a=hello&a=good&a3=xxxx&username=admin&x=333#wechat
	fmt.Println(s)
}
