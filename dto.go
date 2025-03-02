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

type MaxwellDto struct {
	Database          string          `json:"database"` // 库名
	Table             string          `json:"table"`    // 表名
	Type              string          `json:"type"`     // 类型
	Ts                int64           `json:"ts"`       //秒级别的时间戳
	Xid               int64           `json:"xid"`      //事务id
	Commit            bool            `json:"commit"`
	Position          string          `json:"position,omitempty"`
	ServerID          int64           `json:"server_id,omitempty"`
	ThreadID          int64           `json:"thread_id,omitempty"`
	PrimaryKey        []interface{}   `json:"primary_key,omitempty"`
	PrimaryKeyColumns []string        `json:"primary_key_columns,omitempty"`
	Data              json.RawMessage `json:"data"`
	Old               json.RawMessage `json:"old,omitempty"`
}

type IdReqDto struct {
	Id int64 `json:"id" binding:"required"`
}

type IdsReqDto struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type IdRespDto struct {
	Id int64 `json:"id"`
}
