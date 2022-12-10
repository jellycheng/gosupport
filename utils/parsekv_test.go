package utils

import (
	"fmt"
	"testing"
)

//go test -run="TestNewParseKV"
func TestNewParseKV(t *testing.T) {
	//示例1
	obj := NewParseKV()
	obj.SetCommnetFlag("#;").SetContent(" abc='123' ")
	fmt.Println(obj.GetContent())
	obj.Parse()
	key := obj.GetResultKey()
	val := obj.GetResultVal()
	comment := obj.GetResultComment()
	//key=abc,val=123,comment=
	fmt.Println(fmt.Sprintf("key=%s,val=%s,comment=%s", key, val, comment))

	//示例2: key=abc,val=123,comment=我是注释内容了
	obj2 := NewParseKV()
	obj2.SetCommnetFlag("#;").SetContent(" abc='123' #我是注释内容了 ").Parse()
	fmt.Println(fmt.Sprintf("key=%s,val=%s,comment=%s", obj2.GetResultKey(), obj2.GetResultVal(), obj2.GetResultComment()))

	//示例3: key=abc,val='"xyz",comment=我是注释内容了
	obj3 := NewParseKV()
	obj3.SetCommnetFlag("#;").SetContent(` abc="'"xyz"" ;我是注释内容了 `).Parse()
	fmt.Println(fmt.Sprintf("key=%s,val=%s,comment=%s", obj3.GetResultKey(), obj3.GetResultVal(), obj3.GetResultComment()))

	//示例4: key=,val=,comment=abc=注释内容了
	obj4 := NewParseKV()
	obj4.SetCommnetFlag("#").SetContent(` #abc=注释内容了 `).Parse()
	fmt.Println(fmt.Sprintf("key=%s,val=%s,comment=%s", obj4.GetResultKey(), obj4.GetResultVal(), obj4.GetResultComment()))

}
