package curl

import "testing"

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
