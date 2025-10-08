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

func (m *HttpRequest) SetDialTimeOut(TimeOutSecond int64) *HttpRequest {
	m.dialTimeout = time.Duration(TimeOutSecond)
	return m
}

func (m *HttpRequest) SetResponseTimeOut(TimeOutSecond int64) *HttpRequest {
	m.responseTimeOut = time.Duration(TimeOutSecond)
	return m
}

func (m *HttpRequest) SetTimeout(TimeOutSecond int64) *HttpRequest {
	m.timeout = time.Duration(TimeOutSecond)
	return m
}

func (m *HttpRequest) SetPostType(typeStr string) *HttpRequest {
	m.postType = strings.ToLower(typeStr)
	return m
}

func (m *HttpRequest) GetTimeout() time.Duration {
	return m.timeout
}

// 设置请求方法,返回Request结构体对象用于链式调用
func (m *HttpRequest) SetMethod(method string) *HttpRequest {
	m.method = strings.ToUpper(method)
	return m
}

func (m HttpRequest) GetMethod() string {
	return m.method
}

// 设置请求地址
func (m *HttpRequest) SetUrl(url string) *HttpRequest {
	m.url = url
	return m
}

func (m HttpRequest) GetUrl() string {
	return m.url
}

// 设置请求头
func (m *HttpRequest) SetHeaders(headers map[string]string) *HttpRequest {
	m.headers = headers
	return m
}

func (m *HttpRequest) AddHeader(header, val string) *HttpRequest {
	m.headers[header] = val
	return m
}

func (m HttpRequest) GetHeaders() map[string]string {
	return m.headers
}

// 设置请求cookies
func (m *HttpRequest) SetCookies(cookies map[string]string) *HttpRequest {
	m.cookies = cookies
	return m
}

func (m *HttpRequest) GetCookies() map[string]string {
	return m.cookies
}

// 设置url查询参数
func (m *HttpRequest) SetQueries(queries map[string]string) *HttpRequest {
	m.queries = queries
	return m
}

func (m HttpRequest) GetQueries() map[string]string {
	return m.queries
}

// 设置post请求的提交数据
func (m *HttpRequest) SetPostData(postData map[string]interface{}) *HttpRequest {
	m.postData = postData
	return m
}

func (m HttpRequest) GetPostData() map[string]interface{} {
	return m.postData
}

func (m *HttpRequest) SetRawPostData(rawPostData string) *HttpRequest {
	m.rawPostData = rawPostData
	return m
}

func (m HttpRequest) GetRawPostData() string {
	return m.rawPostData
}

// 发起get请求
func (m *HttpRequest) Get() (*HttpResponse, error) {
	m.SetMethod(http.MethodGet)
	return m.send()
}

// 发起post请求
func (m *HttpRequest) Post() (*HttpResponse, error) {
	m.SetMethod(http.MethodPost)
	return m.send()
}

// 发起Put请求
func (m *HttpRequest) Put() (*HttpResponse, error) {
	m.SetMethod(http.MethodPut)
	return m.send()
}

// 发起PATCH请求
func (m *HttpRequest) Patch() (*HttpResponse, error) {
	m.SetMethod(http.MethodPatch)
	return m.send()
}

// 发起Delete请求
func (m *HttpRequest) Delete() (*HttpResponse, error) {
	m.SetMethod(http.MethodDelete)
	return m.send()
}

func (m *HttpRequest) Request() (*HttpResponse, error) {
	r, err := m.send()
	return r, err
}

// 发起请求
func (m *HttpRequest) send() (*HttpResponse, error) {
	url := m.GetUrl()
	if url == "" {
		return nil, errors.New("请求地址不能为空")
	}
	method := m.GetMethod()
	if method == "" {
		return nil, errors.New("请求方式不能为空")
	}

	// 初始化HttpResponse响应对象
	response := NewHttpResponse()

	// 初始化http.Client对象
	m.cli = &http.Client{
		Timeout: m.timeout,
	}

	//处理参数
	var body io.Reader
	/**
	if method == "POST" && m.GetPostData() != nil {
		if jsonData, err := json.Marshal(m.GetPostData()); err != nil {
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
		if m.postType == "json" {
			m.AddHeader("Content-Type", "application/json")
			if rawBody := m.GetRawPostData(); rawBody != "" {
				body = strings.NewReader(rawBody)
			} else if m.GetPostData() != nil {
				if jsonData, err := json.Marshal(m.GetPostData()); err != nil {
					return nil, err
				} else {
					body = bytes.NewReader(jsonData)
				}
			} else {
				body = nil
			}
		} else if m.postType == "form" {
			m.AddHeader("Content-Type", "application/x-www-form-urlencoded")
			if rawBody := m.GetRawPostData(); rawBody != "" {
				body = strings.NewReader(rawBody)
			} else if m.GetPostData() != nil {
				tmpPostData := m.GetPostData()
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
		for k, v := range m.GetQueries() {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		m.req = req
	}

	//处理cookie
	for k, v := range m.GetCookies() {
		m.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	//处理header
	for k, v := range m.GetHeaders() {
		m.req.Header.Set(k, v)
	}

	//发起请求
	if resp, err := m.cli.Do(m.req); err != nil {
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
