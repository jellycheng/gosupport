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
	apisign.AppendParam("a1", nil)
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
	apisign3.AppendParam("username", "tom").AppendParam("age", 18).AppendParam("99", 199).AppendParam("0", 70).AppendParam("fv", float64(123456789012345))
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

//go test -run="TestNewUrlGetSign"
func TestNewUrlGetSign(t *testing.T) {
	s := NewUrlGetSign().SetSecret("cjsJellySecret123456").AutoSpellUrlParam()
	fmt.Println(s)
}

//go test -run="TestDingtalkSign"
func TestDingtalkSign(t *testing.T) {
	secret := "改成机器人密钥"
	timestamp := gosupport.TimeNowMillisecond()
	webhook := "这是机器人webhook地址" //https://oapi.dingtalk.com/robot/send?access_token=访问令牌
	sign := DingtalkSign(timestamp, secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	fmt.Println(url)

}

//go test -run="TestWxCheckSign_Check"
func TestWxCheckSign_Check(t *testing.T) {
	obj := WxCheckSign {
			Token: "helloWorld123!",
			Nonce:"abc123",
			Timestamp:"1648881407",

	}
	sign := WxSignV2(obj.Token,obj.Timestamp, obj.Nonce)
	obj.Signature = sign
	fmt.Println(sign)
	if obj.Check() {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}

}

