package xcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func Sha256V1(str string) string  {
	h := sha256.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256ByByteV1(by []byte) string  {
	h := sha256.New()
	h.Write(by)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSha256V1 HmacSHA256加密算法
func HmacSha256V1(source, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
