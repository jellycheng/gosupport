package sign

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

type WxCheckSign struct {
	Token     string // token，在公众号后台设置的值
	Timestamp string // 时间戳
	Nonce     string // 随机数
	Echostr   string // 随机字符串
	Signature string // 签名串
}

func (m WxCheckSign) Check() bool {
	return WxSignV1(m.Nonce, m.Timestamp, m.Token) == m.Signature
}

func WxSignV1(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))

}

func WxSignV2(token, timestamp, nonce string) string {
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	var str string
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
