package xcrypto

import "encoding/base64"

func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}
