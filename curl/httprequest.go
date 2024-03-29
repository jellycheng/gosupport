package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strings"
	"time"
)

// 请求结构体
type HttpRequest struct {
	cli             *http.Client
	req             *http.Request
	method          string
	url             string
	dialTimeout     time.Duration
	responseTimeOut time.Duration
	timeout         time.Duration //请求超时
	headers         map[string]string
	cookies         map[string]string
	queries         map[string]string
	postData        map[string]interface{}
	rawPostData     string //原始post数据，与postData二选一,rawPostData >postData
	postType        string //post请求方式，form,json，text_plain

}

// 创建一个Request对象
func NewHttpRequest() *HttpRequest {
	r := &HttpRequest{}
	r.dialTimeout = 10 * time.Second
	r.responseTimeOut = 10 * time.Second
	r.timeout = 5 * time.Second
	r.headers = make(map[string]string)
	r.cookies = make(map[string]string)
	r.queries = make(map[string]string)
	r.postData = make(map[string]interface{})
	r.postType = "form"
	return r
}

func (this *HttpRequest) SetDialTimeOut(TimeOutSecond int64) *HttpRequest {
	this.dialTimeout = time.Duration(TimeOutSecond)
	return this
}

func (this *HttpRequest) SetResponseTimeOut(TimeOutSecond int64) *HttpRequest {
	this.responseTimeOut = time.Duration(TimeOutSecond)
	return this
}

func (this *HttpRequest) SetTimeout(TimeOutSecond int64) *HttpRequest {
	this.timeout = time.Duration(TimeOutSecond)
	return this
}

func (this *HttpRequest) SetPostType(typeStr string) *HttpRequest {
	this.postType = strings.ToLower(typeStr)
	return this
}

func (this *HttpRequest) GetTimeout() time.Duration {
	return this.timeout
}

// 设置请求方法,返回Request结构体对象用于链式调用
func (this *HttpRequest) SetMethod(method string) *HttpRequest {
	this.method = strings.ToUpper(method)
	return this
}

func (this HttpRequest) GetMethod() string {
	return this.method
}

// 设置请求地址
func (this *HttpRequest) SetUrl(url string) *HttpRequest {
	this.url = url
	return this
}

func (this HttpRequest) GetUrl() string {
	return this.url
}

// 设置请求头
func (this *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	this.headers = headers
	return this
}

func (this *HttpRequest) AddHeader(header, val string) *HttpRequest {
	this.headers[header] = val
	return this
}

func (this HttpRequest) GetHeaders() map[string]string {
	return this.headers
}

// 设置请求cookies
func (this *HttpRequest) SetCookies(cookies map[string]string) *HttpRequest {
	this.cookies = cookies
	return this
}

func (this *HttpRequest) GetCookies() map[string]string {
	return this.cookies
}

// 设置url查询参数
func (this *HttpRequest) SetQueries(queries map[string]string) *HttpRequest {
	this.queries = queries
	return this
}

func (this HttpRequest) GetQueries() map[string]string {
	return this.queries
}

// 设置post请求的提交数据
func (this *HttpRequest) SetPostData(postData map[string]interface{}) *HttpRequest {
	this.postData = postData
	return this
}

func (this HttpRequest) GetPostData() map[string]interface{} {
	return this.postData
}

func (this *HttpRequest) SetRawPostData(rawPostData string) *HttpRequest {
	this.rawPostData = rawPostData
	return this
}

func (this HttpRequest) GetRawPostData() string {
	return this.rawPostData
}

// 发起get请求
func (this *HttpRequest) Get() (*HttpResponse, error) {
	this.SetMethod(http.MethodGet)
	return this.send()
}

// 发起post请求
func (this *HttpRequest) Post() (*HttpResponse, error) {
	this.SetMethod(http.MethodPost)
	return this.send()
}

// 发起Put请求
func (this *HttpRequest) Put() (*HttpResponse, error) {
	this.SetMethod(http.MethodPut)
	return this.send()
}

// 发起PATCH请求
func (this *HttpRequest) Patch() (*HttpResponse, error) {
	this.SetMethod(http.MethodPatch)
	return this.send()
}

// 发起Delete请求
func (this *HttpRequest) Delete() (*HttpResponse, error) {
	this.SetMethod(http.MethodDelete)
	return this.send()
}

func (this *HttpRequest) Request() (*HttpResponse, error) {
	r, err := this.send()
	return r, err
}

//发起请求
func (this *HttpRequest) send() (*HttpResponse, error) {
	url := this.GetUrl()
	if url == "" {
		return nil, errors.New("请求地址不能为空")
	}
	method := this.GetMethod()
	if method == "" {
		return nil, errors.New("请求方式不能为空")
	}

	// 初始化HttpResponse响应对象
	response := NewHttpResponse()

	// 初始化http.Client对象
	this.cli = &http.Client{
		Timeout: this.timeout,
	}

	//处理参数
	var body io.Reader
	/**
	if method == "POST" && this.GetPostData() != nil {
		if jsonData, err := json.Marshal(this.GetPostData()); err != nil {
			return nil, err
		} else {
			body = bytes.NewReader(jsonData)
		}
	} else {
		body = nil
	}
	*/
	switch method {
	case http.MethodGet, http.MethodDelete:
		body = nil
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions:
		if this.postType == "json" {
			this.AddHeader("Content-Type", "application/json")
			if rawBody := this.GetRawPostData(); rawBody != "" {
				body = strings.NewReader(rawBody)
			} else if this.GetPostData() != nil {
				if jsonData, err := json.Marshal(this.GetPostData()); err != nil {
					return nil, err
				} else {
					body = bytes.NewReader(jsonData)
				}
			} else {
				body = nil
			}
		} else if this.postType == "form" {
			this.AddHeader("Content-Type", "application/x-www-form-urlencoded")
			if rawBody := this.GetRawPostData(); rawBody != "" {
				body = strings.NewReader(rawBody)
			} else if this.GetPostData() != nil {
				tmpPostData := this.GetPostData()
				tmpValues := neturl.Values{}
				for k, v := range tmpPostData {
					if vv, ok := v.(string); ok {
						tmpValues.Set(k, vv)
					} else if vv, ok := v.([]string); ok {
						for _, vvv := range vv {
							tmpValues.Add(k, vvv)
						}
					} else {
						tmpValues.Set(k, fmt.Sprintf("%v", v))
					}
				}
				body = strings.NewReader(tmpValues.Encode())
				//body = bytes.NewReader(jsonData)
			} else {
				body = nil
			}
		} else {
			body = nil
		}
	default:
		return nil, errors.New("无效的请求方式")
	}

	//请求对象
	if req, err := http.NewRequest(method, url, body); err != nil {
		return nil, err
	} else {
		//处理get参数
		q := req.URL.Query()
		for k, v := range this.GetQueries() {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		this.req = req
	}

	//处理cookie
	for k, v := range this.GetCookies() {
		this.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	//处理header
	for k, v := range this.GetHeaders() {
		this.req.Header.Set(k, v)
	}

	//发起请求
	if resp, err := this.cli.Do(this.req); err != nil {
		return nil, err
	} else {
		response.SetRaw(resp)
	}
	defer response.GetRaw().Body.Close()

	//解析响应头
	response.parseHeaders()
	//解析响应body
	response.parseBody()

	return response, nil
}
