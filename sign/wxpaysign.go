package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/jellycheng/gosupport"
	"hash"
	"sort"
	"strings"
)

/*
实现微信支付签名
https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=4_3
signType = MD5 | HMAC-SHA256
PayKey =支付密钥
*/
func WxPaySign(params map[string]string, signType string, PayKey string) (sign, str string, err error) {
	var kvs []string
	for k, v := range params {
		if len(v) > 0 && strings.ToLower(k) != "sign" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(kvs)
	kvs = append(kvs, fmt.Sprintf("key=%s", PayKey))
	str = strings.Join(kvs, "&")

	var h hash.Hash
	if signType == gosupport.SignTypeHmacSHA256 {
		h = hmac.New(sha256.New, []byte(PayKey))
	} else {
		h = md5.New()
	}
	if _, err = h.Write([]byte(str)); err != nil {
		return
	}

	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	return
}
