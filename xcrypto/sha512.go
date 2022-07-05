package xcrypto

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(dataStr string) string {
	h := sha512.New()
	h.Write([]byte(dataStr))
	return hex.EncodeToString(h.Sum([]byte("")))
}
