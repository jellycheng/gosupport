package sign

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"testing"
)


//go test -run="TestNewApiSign01"
func TestNewApiSign01(t *testing.T) {

	apisign := NewApiSign01()
	apisign.SetSecret("abc123$%^cjs789xyz").AppendParam("userid", 123).AppendParam("sign", "123").AppendParam("hello", "")
	apisign.AppendParam("username", "admin").AppendParam("nickname", "你好，我是xxx")
	fmt.Println(apisign.Md5Sign())
	fmt.Println(apisign.GetSignString())

	fmt.Println()

	apisign2 := NewApiSign01()
	apisign2.SetSecret("abc123$%^cjs789xyz")
	fmt.Println(apisign2.Md5Sign())
	fmt.Println(apisign2.GetSignString())

	fmt.Println()

	apisign3 := new(Apisign01)
	apisign3.SetSecret("密钥123")
	//apisign3.SetParams(make(map[string]interface{}))
	//apisign3.SetParams(map[string]interface{}{})
	apisign3.SetParams(map[string]interface{}{"xyz":99})
	apisign3.AppendParam("username", "tom").AppendParam("age", 18).AppendParam("99", 199).AppendParam("0", 70)
	fmt.Println(apisign3.Md5Sign())
	fmt.Println(apisign3.GetSignString())

}

//go test -run="TestWxPaySign"
func TestWxPaySign(t *testing.T) {
	params := map[string]string {
		"xyz": "123",
		"ab": "我是ab",
		"appid": "wx123456789",
	}
	payKey := "dfasdf12312323xdfaaf12323x"
	if sign, str, err := WxPaySign(params, gosupport.SignTypeMD5, payKey);err == nil {
		fmt.Println("str:", str)
		fmt.Println("sign:", sign)
	} else {
		fmt.Println(err.Error())
	}

	if sign, str, err := WxPaySign(params, gosupport.SignTypeHmacSHA256, payKey);err == nil {
		fmt.Println("str:", str)
		fmt.Println("sign:", sign)
	} else {
		fmt.Println(err.Error())
	}

}
