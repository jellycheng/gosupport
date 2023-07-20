package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/jellycheng/gosupport"
	"sort"
	"strings"
)

// XimalayaSign 封装喜马拉雅开放平台签名算法
type XimalayaSign struct {
	appSecret                   string                 //app_secret
	serverAuthenticateStaticKey string                 // 密钥
	params                      map[string]interface{} // 参数
	signStr                     string                 // 拼接原串
}

func (m *XimalayaSign) SetAppSecret(s string) *XimalayaSign {
	m.appSecret = s
	return m
}

func (m *XimalayaSign) SetServerAuthenticateStaticKey(s string) *XimalayaSign {
	m.serverAuthenticateStaticKey = s
	return m
}

func (m *XimalayaSign) SetParams(p map[string]interface{}) *XimalayaSign {
	if m.params == nil {
		m.params = make(map[string]interface{})
	}
	m.params = p
	return m
}

func (m *XimalayaSign) AppendParam(k string, v interface{}) *XimalayaSign {
	m.params[k] = v
	return m
}

func (m *XimalayaSign) GetSign() string {
	sign := ""
	var keys []string
	for k, _ := range m.params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	if len(keys) != 0 {
		for _, v := range keys {
			builder.WriteString(fmt.Sprintf("%s=%s&", v, gosupport.ToStr(m.params[v])))
		}
	}
	tmpStr := builder.String()
	rs := []rune(tmpStr)
	l := len(rs)
	if l == 0 {
		m.signStr = ""
	} else {
		m.signStr = string(rs[0 : l-1])
	}
	s1 := base64.StdEncoding.EncodeToString([]byte(m.signStr))
	sha1Key := m.appSecret + m.serverAuthenticateStaticKey
	s2 := m.HmacSha1(s1, sha1Key)
	sign = gosupport.Md5V1(string(s2))

	return sign
}

func (m XimalayaSign) HmacSha1(dataStr, key string) []byte {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(dataStr))
	return h.Sum([]byte(""))
}

func (m XimalayaSign) GetSignSrc() string {
	return m.signStr
}

func NewXimalayaSign() *XimalayaSign {
	signObj := new(XimalayaSign)
	signObj.params = make(map[string]interface{})
	return signObj
}
