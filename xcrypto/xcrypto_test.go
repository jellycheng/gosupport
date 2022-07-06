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

// go test -run=TestAesCbcDecrypt
func TestAesCbcDecrypt(t *testing.T) {
	encrypted := []byte(Base64StdDecode("+Lh80bwokENX/RQWZdZtaSyQBC15k5P7eszF60HEe0q4wy4Ct1wlJJBpFG25VnCexql9nqMJSbr2PdWXJKBC/O4wAE6/jlOztbWGJ+sMu8hqub/DxX3Cv2Lv9IOAbuyIQnpKF6ZFKuGS3vJFmFhEttMX8fszO2lsksr/bLKJFwc="))
	key := []byte("85k8HJF3wqy1YSaG")
	res := AesCbcDecrypt(encrypted, key)
	fmt.Println(string(res))

}

// go test -run=TestAesCbcEncrypt
func TestAesCbcEncrypt(t *testing.T) {
	str:=`
{
	"appId":"abc123",
	"userId":"xxxtest"
}
`
	k := "85k8HJF3wqy1YSaGNk1a7J9pctE2gONR"
	// 加密
	bs := AesCbcEncrypt([]byte(str), []byte(k))
	bsStr := Base64StdEncode(string(bs))
	fmt.Println(bsStr)

	// 解密
	bs2 := AesCbcDecrypt([]byte(Base64StdDecode(bsStr)), []byte(k))
	fmt.Println(string(bs2))
	//
	bs3 := AesCbcDecrypt(bs, []byte(k))
	fmt.Println(string(bs3))
}
