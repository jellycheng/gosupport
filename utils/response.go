package utils

type EmptyStruct struct{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	//omitempty表值如果为空则忽略该字段，-表直接忽略该字段，如：abc string `json:"-"`
	Trace_id string `json:"trace_id,omitempty"`
}

func NewResponse() *Response {
	res := &Response{Code: 0, Data: struct{}{}}
	return res
}

func Success() *Response {
	res := &Response{
		Code: 0,
		Data: struct{}{},
		Msg:  "OK",
	}
	return res
}

func Fail() *Response {
	res := &Response{
		Code: 1000,
		Data: struct {
		}{},
		Msg: "fail",
	}
	return res
}

func (res *Response) SetCode(code int) *Response {
	res.Code = code
	return res
}
func (res *Response) GetCode() int {
	return res.Code
}

func (res *Response) SetMsg(msg string) *Response {
	res.Msg = msg
	return res
}
func (res *Response) GetMsg() string {
	return res.Msg
}

func (res *Response) SetData(data interface{}) *Response {
	res.Data = data
	return res
}

func (res *Response) GetData() interface{} {
	return res.Data
}

func (res *Response) SetTraceid(trace_id string) *Response {
	res.Trace_id = trace_id
	return res
}
func (res *Response) GetTraceid() string {
	return res.Trace_id
}

// 公共分页列表数据结构
type CommonListRespDto struct {
	Page     int         `json:"page"`      //第几页
	PageSize int         `json:"page_size"` //每页显示几条
	Total    int64       `json:"total"`     //总记录数
	List     interface{} `json:"list"`
}

func NewCommonListRespDto() CommonListRespDto {
	dto := CommonListRespDto{}
	return dto
}

func NewCommonListRespDtoObj() *CommonListRespDto {
	dto := &CommonListRespDto{}
	return dto
}
