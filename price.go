package gosupport

import (
	"fmt"
	"strings"
)

// 分转元
func Fen2yuan4int(price int, isTrimZero bool) string {
	return Fen2yuan(int64(price), isTrimZero)
}

/*
	分转元
	price := 1909
	s1 := gosupport.Fen2yuan(int64(price), true)
*/
func Fen2yuan(price int64, isTrimZero bool) string {
	mathSymbol := ""
	if price < 0 {
		mathSymbol = "-"
		price = price * -1
	}
	tmpPrice01 := price / 100
	//tmpPrice02 := price - tmpPrice01 * 100
	tmpPrice02 := price % 100
	formatPrice := ""
	if tmpPrice02 < 10 && tmpPrice02 > 0 {
		formatPrice = fmt.Sprintf("%s%d.0%d", mathSymbol, tmpPrice01, tmpPrice02)
	} else if tmpPrice02 == 0 {
		if isTrimZero {
			isTrimZero = false
			formatPrice = fmt.Sprintf("%s%d", mathSymbol, tmpPrice01)
		} else {
			formatPrice = fmt.Sprintf("%s%d.00", mathSymbol, tmpPrice01)
		}
	} else {
		formatPrice = fmt.Sprintf("%s%d.%d", mathSymbol, tmpPrice01, tmpPrice02)
	}

	if isTrimZero {
		formatPrice = strings.TrimRight(formatPrice, "0")
	}

	return formatPrice
}

// 调用示例：gosupport.RemoveDian00("18.00") 返回 18
func RemoveDian00(priceStr string) string {
	if priceStr == "" {
		return priceStr
	}
	parts := strings.Split(priceStr, ".")
	if len(parts) == 2 {
		tmpStr := strings.TrimRight(parts[1], ".0")
		if tmpStr == "" {
			return parts[0]
		} else {
			return parts[0] + "." + tmpStr
		}
	} else {
		tmpStr := strings.TrimRight(priceStr, ".0")
		if tmpStr != "" {
			return tmpStr
		}
		return "0"
	}
}

// GiveUpFloat64ToInt64 舍弃小数取整
func GiveUpFloat64ToInt64(f float64) int64 {
	return int64(f)
}
