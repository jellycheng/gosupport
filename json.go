package gosupport

import "encoding/json"

// ToJson 转成json字符串
func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JsonUnmarshal(str string, obj interface{}) error  {
	return json.Unmarshal([]byte(str), obj)
}
