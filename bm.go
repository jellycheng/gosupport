package gosupport

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type BodyMap map[string]interface{}

func (m BodyMap) Set(key string, value interface{}) BodyMap {
	m[key] = value
	return m
}

func (m BodyMap) GetInterface(key string) interface{} {
	if m == nil {
		return nil
	}
	v, ok := m[key]
	if !ok {
		return nil
	}
	return v
}

func (m BodyMap) GetString(key string) string {
	s := m.GetInterface(key)
	if s == nil {
		return ""
	}
	return ToStr(s)
}

func (m BodyMap) Remove(key string) {
	delete(m, key)
}

func (m BodyMap) Reset() {
	for k := range m {
		delete(m, k)
	}
}

func (m BodyMap) AddSubBodyMap(key string, fn func(sub BodyMap)) BodyMap {
	_newMObj := make(BodyMap)
	fn(_newMObj)
	m[key] = _newMObj
	return m
}

func (m BodyMap) ToJson() string {
	return ToJson(m)
}

func (m BodyMap) CheckEmpty(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := m.GetString(k); v == "" {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return fmt.Errorf("%s： %v", "缺少必需的参数", strings.Join(emptyKeys, ", "))
	}
	return nil
}

func (m BodyMap) UrlEncode() string {
	if m == nil {
		return ""
	}
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := m.GetString(k); v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

func NewBodyMap() BodyMap {
	m := make(BodyMap)
	return m
}
