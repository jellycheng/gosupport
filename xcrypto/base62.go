package xcrypto

import (
	"bytes"
	"math"
)

const ShortChar = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(number int64) string {
	if number == 0 {
		return string(ShortChar[0])
	}
	chars := make([]byte, 0)
	length := len(ShortChar) //62

	for number > 0 {
		result := number / int64(length)
		remainder := number % int64(length) //余数
		chars = append(chars, ShortChar[remainder])
		number = result
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func Base62Decode(str string) int64 {
	var number int64 = 0
	idx := 0.0
	chars := []byte(ShortChar)
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
