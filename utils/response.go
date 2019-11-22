package utils

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	//omitempty表值如果为空则忽略该字段，-表直接忽略该字段，如：abc string `json:"-"`
	Trace_id   string      `json:"trace_id,omitempty"`
}

func Success() *Response {
	res := &Response{
		Code:  0,
		Data: struct { }{},
		Msg:   "OK",
	}
	return res
}

func Fail() *Response {
	res := &Response{
		Code:  1000,
		Data: struct {

		}{},
		Msg:   "fail",
	}
	return res
}

func (res *Response) SetCode(code int)  {
	res.Code = code
}
func (res *Response) GetCode() int  {
	return res.Code
}

func (res *Response) SetMsg(msg string)  {
	res.Msg = msg
}
func (res *Response) GetMsg() string  {
	return res.Msg
}

func (res *Response) SetData(data interface{})  {
	res.Data = data
}

func (res *Response) GetData() interface{}  {
	return res.Data
}

func (res *Response) SetTraceid(trace_id string)  {
	res.Trace_id = trace_id
}
func (res *Response) GetTraceid() string  {
	return res.Trace_id
}
