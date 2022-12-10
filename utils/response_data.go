package utils

import "encoding/json"

// 对data数据进行增删改查

type ResponseData struct {
	Data map[string]interface{}
}

func NewResponseData() *ResponseData {
	return &ResponseData{
		Data: make(map[string]interface{}),
	}
}

// 增、改
func (builder *ResponseData) Put(key string, value interface{}) *ResponseData {
	builder.Data[key] = value
	return builder
}

// 获取
func (builder *ResponseData) Get(key string) (value interface{}, exists bool) {
	value, exists = builder.Data[key]
	return
}

//删
func (builder *ResponseData) Del(key string) *ResponseData {
	delete(builder.Data, key)
	return builder
}

func (builder *ResponseData) Build() map[string]interface{} {
	return builder.Data
}

func (builder *ResponseData) ToJson() string {
	b, err := json.Marshal(builder.Data)
	if err != nil {
		return ""
	}
	return string(b)
}
