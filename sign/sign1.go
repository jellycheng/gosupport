package sign

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"sort"
	"strings"
)

// Apisign01 签名算法1
type Apisign01 struct {
	secret  string
	params  map[string]interface{}
	signStr string
}

func (m *Apisign01) SetSecret(s string) *Apisign01 {
	m.secret = s
	return m
}

func (m *Apisign01) SetParams(p map[string]interface{}) *Apisign01 {
	if m.params == nil {
		m.params = make(map[string]interface{})
	}
	m.params = p
	return m
}

func (m *Apisign01) AppendParam(k string, v interface{}) *Apisign01 {
	m.params[k] = v
	return m
}

func (m *Apisign01) Md5Sign() string {
	sign := ""
	var keys []string
	for k, v := range m.params {
		if k == "sign" || v == "" || v == nil {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	if len(keys) != 0 {
		for _, v := range keys {
			//builder.WriteString(v)
			//builder.WriteString("=")
			//builder.WriteString(fmt.Sprint(m.params[v]))
			//builder.WriteString("&")
			builder.WriteString(fmt.Sprintf("%s=%s&", v, gosupport.ToStr(m.params[v])))

		}
		builder.WriteString("key=" + m.secret)
	} else {
		builder.WriteString("&key=" + m.secret)
	}
	m.signStr = builder.String()
	sign = strings.ToUpper(gosupport.Md5V1(m.signStr))

	return sign
}

func (m Apisign01) GetSignString() string  {
	return m.signStr
}

func NewApiSign01() *Apisign01 {
	apisign := new(Apisign01)
	apisign.params = make(map[string]interface{})
	return apisign
}



