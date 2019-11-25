package curl

import (
	"fmt"
	"testing"
)

func TestHttpRequest_Get(t *testing.T) {
	req1 := NewHttpRequest()

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
		"interests": []string{"篮球", "旅游", "听音乐"},
		"isAdmin":   true,
	}
	res2, err := req2.SetUrl("http://devapi.nfangbian.com/test.php?a=2&b=say123").SetPostData(postData).SetQueries(queries).SetCookies(cookies).SetHeaders(headers).Post()
	if err==nil {
		fmt.Println("响应结果2：", res2.GetBody())
	} else {
		fmt.Println(err)
	}

}

