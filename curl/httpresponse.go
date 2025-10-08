package curl

import (
	"io"
	"net/http"
)

type HttpResponse struct {
	raw     *http.Response
	headers map[string]string
	body    string
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{}
}

func (m *HttpResponse) SetRaw(r *http.Response) *HttpResponse {
	m.raw = r
	return m
}

func (m *HttpResponse) GetRaw() *http.Response {
	return m.raw
}

func (m *HttpResponse) IsOk() bool {
	return m.raw.StatusCode == 200
}

func (m *HttpResponse) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range m.raw.Header {
		headers[k] = v[0]
	}
	m.headers = headers
	return nil
}

func (m HttpResponse) GetHeaders() map[string]string {
	return m.headers
}

func (m *HttpResponse) parseBody() error {
	if body, err := io.ReadAll(m.raw.Body); err != nil {
		//发生错误
		panic(err)
	} else {
		m.body = string(body)
	}
	return nil
}

// 响应内容
func (m HttpResponse) GetBody() string {
	return m.body
}
