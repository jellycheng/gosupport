package sign

import (
	"encoding/hex"
	"strings"
)

// StringToHex 字符串转16进制
func StringToHex(s string) string {
	return hex.EncodeToString([]byte(s))
}

// HexToString 16进制转字符串
func HexToString(s string) (string, error) {
	s = strings.ReplaceAll(s, " ", "")
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
