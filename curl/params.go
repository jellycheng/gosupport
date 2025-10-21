package curl

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// 自定义类型
type Params map[string]interface{}

// 为类型添加方法，设置参数值
func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

// 批量设置参数
func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

// 获取
func (p Params) GetParam(key string, defaultVal interface{}) interface{} {
	ret, ok := p[key]
	if !ok {
		ret = defaultVal
	}
	return ret
}

// 删除
func (p Params) Del(key string) {
	delete(p, key)
}

// 返回结果示例：empty=&is_supper=true&price=100.25&userid=123&username=tom
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

// ParseQueryParams("app_id=1128&userid=456").Get("app_id")
func ParseQueryParams(params string) url.Values {
	urlValues, err := url.ParseQuery(params)
	if err != nil {
		return url.Values{}
	}
	return urlValues
}

// 分析url并追加参数
func ParseUrlAndAppendParam(reqUrl string, queryParams map[string][]string) (string, error) {
	varURL, err := url.Parse(reqUrl)
	if err != nil {
		return "", err
	}
	query := varURL.Query()
	for k, values := range queryParams {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	varURL.RawQuery = query.Encode()
	return varURL.String(), nil
}
