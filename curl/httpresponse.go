package curl

import (
	"io/ioutil"
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

func (this *HttpResponse) SetRaw(r *http.Response) *HttpResponse {
	this.raw = r
	return this
}

func (this *HttpResponse) GetRaw()  *http.Response {
	return this.raw
}

func (this *HttpResponse) IsOk() bool {
	return this.raw.StatusCode == 200
}

func (this *HttpResponse) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range this.raw.Header {
		headers[k] = v[0]
	}
	this.headers = headers
	return nil
}

func (this HttpResponse) GetHeaders() map[string]string {
	return this.headers
}

func (this *HttpResponse) parseBody() error {
	if body, err := ioutil.ReadAll(this.raw.Body); err != nil {
		//发生错误
		panic(err)
	} else {
		this.body = string(body)
	}
	return nil
}

//响应内容
func (this HttpResponse) GetBody() string  {
	return this.body
}

