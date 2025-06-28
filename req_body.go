package gosupport

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type reqBody struct {
	httpReq *http.Request
}

func NewReqBody(r *http.Request) *reqBody {
	return &reqBody{httpReq: r}
}

func (m *reqBody) ReadBodyContent2Byte() ([]byte, error) {
	var body, err = ioutil.ReadAll(m.httpReq.Body)
	_ = m.httpReq.Body.Close()
	if err != nil {
		return body, err
	}
	m.httpReq.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, err
}

func (m *reqBody) UnMarshalResponse(resp interface{}) error {
	var body, err = m.ReadBodyContent2Byte()
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}
