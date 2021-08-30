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

	for k,v := range pObj {
		println(k, ParamValueToString(v))
	}

}

// go test -run="TestAppendParamToUrl"
func TestAppendParamToUrl(t *testing.T) {
	fmt.Println(AppendParamToUrl("http://nfangbian.com:80/fangan/index?xyz=88#hello", "a=你好&b=goods"))
}
