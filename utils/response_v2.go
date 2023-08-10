package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ResponseV2 struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
	TraceID string          `json:"trace_id,omitempty"`
}

func (r *ResponseV2) Resolve(result interface{}) error {
	if r.Code != 0 {
		return errors.New(r.Msg)
	}

	if err := json.Unmarshal(r.Data, result); err != nil {
		return errors.New(fmt.Sprintf("can't decode data:%s", err.Error()))
	}

	return nil
}

func (r *ResponseV2) ToJson() string {
	if b, err := json.Marshal(r); err == nil {
		return string(b)
	}
	return ""
}
