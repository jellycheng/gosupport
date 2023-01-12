package gosupport

import "encoding/json"

type ListRespDto struct {
	List interface{} `json:"list"`
}

// PagerRespDto 分页
type PagerRespDto struct {
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

type ApiUnRespDto struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
	TraceId string          `json:"trace_id,omitempty"`
}
