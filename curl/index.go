package curl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HttpGet 发起get请求
func HttpGet(urlStr string) (string, error) {
	var content string
	resp, err := http.Get(urlStr)
	if err!=nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	content = string(body)

	return content, nil
}

// HttpPost 发起表单post请求
// param := "name=李四&age=19"
func HttpPost(urlStr string, param string) (string, error)  {
	resp, err := http.Post(urlStr, CONTENT_TYPE_FORM, strings.NewReader(param))
	if err!=nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	content := string(body)

	return content, nil
}

func HttpPostJson(urlStr string, param string) (string, error)  {
	resp, err := http.Post(urlStr, CONTENT_TYPE_JSON, strings.NewReader(param))
	if err!=nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	content := string(body)

	return content, nil
}

// HttpPostForm param := url.Values{"参数名1": {"值1"}, "参数名2": {"值1","值N"}}
func HttpPostForm(urlStr string, param url.Values) (string, error)  {
	resp, err := http.PostForm(urlStr, param)
	if err!=nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	var body, _ = ioutil.ReadAll(resp.Body)
	var content = string(body)
	return content, nil
}

func HttpPostXml(urlStr string, xmlStr string) (string, error)  {
	payload := strings.NewReader(xmlStr)
	req, _ := http.NewRequest("POST", urlStr, payload)
	req.Header.Add("content-type", CONTENT_TYPE_XML)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func ReadBodyContent(httpResp *http.Response) (string, error) {
	var body, err = ReadBodyContent2Byte(httpResp)
	return string(body), err
}

func ReadBodyContent2Byte(httpResp *http.Response) ([]byte, error) {
	var body, err = ioutil.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()
	if err != nil {
		return body, err
	}
	httpResp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, err
}

func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	var body, err = ReadBodyContent2Byte(httpResp)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}
