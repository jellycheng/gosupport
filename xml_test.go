package gosupport

import (
	"fmt"
	"testing"
)

//go test -run="TestMap2XML"
func TestMap2XML(t *testing.T) {
	param := map[string]string{
		"abc": "hello",
		"say": "world'!yes",
	}
	if s, err := Map2XML(param); err == nil {
		fmt.Println(string(s))
	} else {
		fmt.Println(err.Error())
	}
}

// go test -run="TestXML2Map"
func TestXML2Map(t *testing.T) {
	xmlByte := []byte("<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>")

	if res, err := XML2Map(xmlByte);err == nil {
		fmt.Println(fmt.Sprintf("%T, %#v", res, res))

	} else {
		fmt.Println("xml解析错误：", err.Error())
	}

}
