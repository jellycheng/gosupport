package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func BodyContentMd5Sign(body string, secret string) string {
	sign := ""
	str := fmt.Sprintf("%s\n%s", body, secret)
	h := md5.New()
	h.Write([]byte(str))
	sign = hex.EncodeToString(h.Sum(nil))
	return sign
}
