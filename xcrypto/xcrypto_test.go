package xcrypto

import (
	"fmt"
	"testing"
)

// go test -run=TestCreateRsaKey
func TestCreateRsaKey(t *testing.T) {
	_ = CreateRsaKey(4096, "/data/www/cert/rsa_private.pem", "/data/www/cert/rsa_pub.pem")
}

// go test -run=TestRsaEncrypt
func TestRsaEncrypt(t *testing.T) {
	str:=`
{
	"appId":"abc123",
	"userId":"xxxtest"
}
`
	// 加密
	err, d := RsaEncrypt([]byte(str), "/data/www/cert/rsa_pub.pem")
	if err == nil{
		con := Base64StdEncode(string(d))
		fmt.Println(con)
	} else {
		fmt.Println(err.Error())
	}

	// 解密
	if err, d2 := RsaDecrypt(d, "/data/www/cert/rsa_private.pem");err == nil {
		fmt.Println(string(d2))
	} else {
		fmt.Println(err.Error())
	}

}

