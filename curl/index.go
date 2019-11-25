package curl

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
 * 发起get请求
 */
func HttpGet(url string) (string, error) {
	var content string
	resp, err := http.Get(url)
	if err!=nil {
		return "", err
	}
	defer resp.Body.Close()  //关闭资源
	body, _ := ioutil.ReadAll(resp.Body)
	content = string(body)

	return content, nil
}

/**
 * 发起表单post请求
 * param := "name=李四&age=19"
 */
func HttpPost(url string, param string) (string, error)  {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(param))
	if err!=nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	content := string(body)

	return content, nil
}

func HttpPostJson(url string, param string) (string, error)  {
	resp, err := http.Post(url, "application/json", strings.NewReader(param))
	if err!=nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	content := string(body)

	return content, nil
}

/**
 * param := url.Values{"参数名1": {"值1"}, "参数名2": {"值1","值N"}}
 */
func HttpPostForm(url string, param url.Values) (string, error)  {
	resp, err := http.PostForm(url, param)
	if err!=nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	content := string(body)

	return content, nil
}

