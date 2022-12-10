package curl

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//参数管理

//自定义类型
type Params map[string]interface{}

//为类型添加方法，设置参数值
func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

//批量设置参数
func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

//获取
func (p Params) GetParam(key string, defaultVal interface{}) interface{} {
	ret, ok := p[key]
	if !ok {
		ret = defaultVal
	}
	return ret
}

//删除
func (p Params) Del(key string) {
	delete(p, key)
}

//返回结果示例：empty=&is_supper=true&price=100.25&userid=123&username=tom
func (p Params) ToQueryString() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, ParamValueToString(v))
	}
	return u.Encode()
}

func NewParams() Params {
	p := make(Params)
	return p
}

func ParamValueToString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return ""
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}
