package utils

import (
	"bytes"
	"math"
)

type MyintCompress struct {
	CharStr string
}

// s := utils.NewMintCompress().Compress(789)
func NewMintCompress() MyintCompress {
	str := "abcdefghkmnpqrstwxyABCDEFGHKMNPQRSTWXY3456789"
	compObj := MyintCompress{
		CharStr: str,
	}
	return compObj
}

func (m MyintCompress) Compress(number int64) string {
	if number == 0 {
		return string(m.CharStr[0])
	}
	chars := make([]byte, 0)
	length := len(m.CharStr)

	for number > 0 {
		result := number / int64(length)
		remainder := number % int64(length)
		chars = append(chars, m.CharStr[remainder])
		number = result
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func (m MyintCompress) UnCompress(str string) int64 {
	var number int64 = 0
	idx := 0.0
	chars := []byte(m.CharStr)
	charsLength := float64(len(chars))
	tokenLength := float64(len(str))
	for _, c := range []byte(str) {
		power := tokenLength - (idx + 1)
		index := bytes.IndexByte(chars, c)
		number += int64(index) * int64(math.Pow(charsLength, power))
		idx++
	}
	return number
}
