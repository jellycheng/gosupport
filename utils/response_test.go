package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {

	fmt.Printf("%#v \n", Success()) //&utils.Response{Code:0, Data:true, Msg:"OK"}

	if json1Str, err := json.Marshal(Success()); err == nil {
		fmt.Println(string(json1Str)) //{"code":0,"data":{},"msg":"OK"}
	} else {
		t.Error("json序列化错误1")
	}

	if json2Str, err := json.Marshal(Fail()); err == nil {
		fmt.Println(string(json2Str)) //{"code":1000,"data":{},"msg":"fail"}
	} else {
		t.Error("json序列化错误2")
	}

	//实例化结构体
	res1 := &Response{}
	res1.SetCode(500)
	res1.SetMsg("发生网络错误了")
	json3Str, err := json.Marshal(res1)
	if err == nil {
		fmt.Println(string(json3Str)) //{"code":500,"data":null,"msg":"发生网络错误了"}
	} else {
		t.Error("json序列化错误3")
	}

	//实例化结构体
	res2 := &Response{Code: 0}
	res2.SetData("成功了")
	res2.SetTraceid("016e922a99e40800274aa49d17950000")
	json4Str, err := json.Marshal(res2)
	if err == nil {
		fmt.Println(string(json4Str)) //{"code":0,"data":"成功了","msg":"","trace_id":"016e922a99e40800274aa49d17950000"}
	} else {
		t.Error("json序列化错误4")
	}

	//=======================注意：要进行json化结构体的字段名首字母必须大写，否则json函数解析不到，如果解析不到则是{}空对象
	type userinfo struct {
		Userid   int    `json:"userid"`
		Nickname string `json:"nickname"`
		Age      int    `json:"age"`
	}
	res3 := &Response{Code: 0}
	res3.SetData(&userinfo{1021, "张三", 18})
	json5Str, err := json.Marshal(res3)
	if err == nil {
		fmt.Println(string(json5Str)) //{"code":0,"data":{"userid":1021,"nickname":"张三","age":18},"msg":""}
	} else {
		t.Error("json序列化错误5")
	}
	fmt.Println(res3.GetData()) //&{1021 张三 18}
	if json6Str, err := json.Marshal(res3.GetData()); err == nil {
		fmt.Println(string(json6Str)) //{"userid":1021,"nickname":"张三","age":18}
	}

	//========
	res4 := &Response{Code: 0}
	res4.SetData(&userinfo{123, "亚瑟", 19}).SetCode(0).SetMsg("ok")
	if json6Str, err := json.Marshal(res4); err == nil {
		//{"code":0,"data":{"userid":123,"nickname":"亚瑟","age":19},"msg":"ok"}
		fmt.Println(string(json6Str))
	} else {
		t.Error("json序列化错误6")
	}

}

// go test -run="TestResponseData"
func TestResponseData(t *testing.T) {
	respData := NewResponseData()
	respData.Put("hello", "world").Put("userid", 123456)
	respData.Put("username", "admin").Put("pwd", "123")
	respData.Put("username", "superadmin")
	respData.Del("pwd").Del("abc")
	fmt.Println(respData.Build())
	fmt.Println(respData.ToJson())
}
