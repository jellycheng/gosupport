package curl

import (
	"fmt"
	"testing"
	"time"
)

func TestHttpRequest_Get(t *testing.T) {
	req1 := NewHttpRequest()
	fmt.Println("默认超时(单位纳秒)为" , int64(req1.GetTimeout()))
	//设置超时
	req1.SetTimeout(int64(15 * time.Second))
	res1, err := req1.SetUrl("http://devapi.nfangbian.com/test.php?a=1&b=hi123").Get()
	if err==nil {
		fmt.Println("响应结果1：", res1.GetBody())
	} else {
		fmt.Println(err)
	}

	//post示例
	req2 := NewHttpRequest()
	headers := map[string]string{
		"User-Agent":    "goland",
		"Authorization": "Bearer access_token 1234",
		"Content-Type":  "application/json",
	}

	cookies := map[string]string{
		"userId":    "1234",
		"loginTime": "15045692188",
	}
	queries := map[string]string{
					"page": "1",
					"pagesize":  "100",
				}
	postData := map[string]interface{}{
		"name":      "admin",
		"age":       24,
		"interests[]": []string{"篮球", "旅游", "听音乐"},
		"isAdmin":   true,
	}
	res2, err := req2.SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetPostData(postData).SetQueries(queries).SetCookies(cookies).SetHeaders(headers).Post()
	if err==nil {
		fmt.Println("响应结果2：", res2.GetBody())
	} else {
		fmt.Println(err)
	}
	fmt.Println("超时设置(单位纳秒)为" , int64(req1.GetTimeout()))
}


func TestHttpRequest02(t *testing.T) {
	//不存在的请求方式
	req2 := NewHttpRequest()
	res2,err := req2.SetMethod("xyz").SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").Request()
	if err==nil {
		fmt.Println("响应结果：", res2.GetBody())
	} else {
		fmt.Println(err.Error())
	}

}

func TestHttpRequest03(t *testing.T) {
	//post-json请求
	req1 := NewHttpRequest()
	postData := map[string]interface{}{
		"name":      "admin123",
		"age":       26,
		"interests": []string{"篮球", "旅游", "听音乐"},
		"isAdmin":   true,
	}
	res1,err := req1.SetMethod("post").SetPostType("json").SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetPostData(postData).Request()
	if err==nil {
		fmt.Println("响应结果1：", res1.GetBody())
	} else {
		fmt.Println(err.Error())
	}

	strJson := "{\"age\":26,\"name\":\"admin123\"}"
	req2 := NewHttpRequest()
	res2,err := req2.SetMethod("post").SetPostType("json").SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetRawPostData(strJson).Request()
	if err==nil {
		fmt.Println("响应结果2：", res2.GetBody())
	} else {
		fmt.Println(err.Error())
	}
}

func TestHttpRequest04(t *testing.T) {
	//post-form请求
	req1 := NewHttpRequest()
	postData := map[string]interface{}{
		"name":      "admin88",
		"age":       18,
		"interests[]": []string{"篮球", "旅游", "听音乐"},
		"isAdmin":   true,
	}
	res1,err := req1.SetMethod("post").SetPostType("form").SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetPostData(postData).Request()
	if err==nil {
		fmt.Println("响应结果1：", res1.GetBody())
	} else {
		fmt.Println(err.Error())
	}

	req2 := NewHttpRequest()
	postDataStr := "xyz=123&username=admin9999"
	res2,err := req2.SetMethod("post").SetPostType("form").SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetRawPostData(postDataStr).Request()
	if err==nil {
		fmt.Println("响应结果2：", res2.GetBody())
	} else {
		fmt.Println(err.Error())
	}
}
