package sign

import (
	"encoding/base64"
	"fmt"
	"github.com/jellycheng/gosupport/xcrypto"
)

// DingtalkSign 钉钉机器人签名
func DingtalkSign(t int64, secret string) string {
	ret := ""
	str := fmt.Sprintf("%d\n%s", t, secret)
	ret = base64.StdEncoding.EncodeToString(xcrypto.HmacSha2562Byte(str, secret))
	return ret
}
