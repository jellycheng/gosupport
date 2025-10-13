package gosupport

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
)

// 去掉所有空格
func TrimAll(src string) string {
	if IsEmptyV2(src) {
		return ""
	}
	src = strings.ReplaceAll(src, "\r", "")
	src = strings.ReplaceAll(src, "\n", "")
	return strings.ReplaceAll(src, " ", "")
}

// hex 转 []byte
func HexToBytes(hexStr string) []byte {
	hexStr = TrimAll(hexStr) // 去掉所有空格
	if IsEmptyV2(hexStr) {
		return []byte("")
	}
	// 解码 hex
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		// 如果输入不是合法 hex，就返回空 slice
		return []byte{}
	}
	return bytes
}

// hex 转 string
func HexToString(hexStr string) string {
	b1 := HexToBytes(hexStr)
	return string(b1)
}

// []byte 转 hex 字符串（不带 0x，全部小写）
func BytesToHex(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return string(dst)
}

// 字符串 转 hex 字符串（不带 0x，全部小写）
func String2Hex(str string) string {
	strBytes := []byte(str)
	hexStr := hex.EncodeToString(strBytes)
	return hexStr
}

func String2HexV2(str string) string {
	str = TrimAll(str)
	return String2Hex(str)
}

func Hex2Base64(hexStr string) string {
	hexBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return ""
	}
	// 将字节数组编码为 Base64 字符串
	base64Str := base64.StdEncoding.EncodeToString(hexBytes)
	return base64Str
}

// aInt64 := gosupport.Hex2int64("0x4090500") 返回 67699968
func Hex2int64(hexStr string) int64 {
	baseVal := 16
	if strings.HasPrefix(hexStr, "0x") || strings.HasPrefix(hexStr, "0X") {
		//hexStr2 := hexStr[2:] // 去掉前缀
		//fmt.Println(hexStr2)
		baseVal = 0
	}
	aInt64, _ := strconv.ParseInt(hexStr, baseVal, 64)
	return aInt64
}

func Base642Hex(src string) string {
	base64Bytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	// 将字节数组编码为十六进制字符串
	hexStr := hex.EncodeToString(base64Bytes)
	return hexStr
}
