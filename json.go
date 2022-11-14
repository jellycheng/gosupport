package gosupport

import (
	"bytes"
	"encoding/json"
)

// ToJson 转成json字符串
func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JsonUnmarshal(str string, obj interface{}) error {
	return json.Unmarshal([]byte(str), obj)
}

// 把变量值转为json字符串且<转为\u003c，>转为\u003e，&转为\u0026
func Escape2json(v interface{}) ([]byte, error) {
	writer := bytes.Buffer{}
	encoder := json.NewEncoder(&writer)
	encoder.SetEscapeHTML(true)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}

	b := writer.Bytes()
	return b[:len(b)-1], nil
}

// 把变量值转为json字符串且\u003c转为<，\u003e转为>，\u0026转为&
func Unescape2json(v interface{}) ([]byte, error) {
	writer := bytes.Buffer{}
	encoder := json.NewEncoder(&writer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}

	b := writer.Bytes()
	return b[:len(b)-1], nil
}
