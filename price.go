package gosupport

import (
	"fmt"
	"strings"
)

/**
	分转元
	price := 1909
	s1 := gosupport.Fen2yuan(int64(price), true)
 */
func Fen2yuan(price int64, isTrimZero bool) string {
	tmpPrice01 := price/100
	//tmpPrice02 := price - tmpPrice01 * 100
	tmpPrice02 := price%100
	formatPrice := ""
	if tmpPrice02<10 && tmpPrice02 >0 {
		formatPrice = fmt.Sprintf("%d.0%d", tmpPrice01, tmpPrice02)
	} else if tmpPrice02 == 0 {
		if isTrimZero {
			isTrimZero = false
			formatPrice = fmt.Sprintf("%d", tmpPrice01)
		} else {
			formatPrice = fmt.Sprintf("%d.00", tmpPrice01)
		}
	} else {
		formatPrice = fmt.Sprintf("%d.%d", tmpPrice01, tmpPrice02)
	}

	if isTrimZero {
		formatPrice = strings.TrimRight(formatPrice, "0")
	}

	return formatPrice
}

// gosupport.RemoveDian00("18.00") 返回 18
func RemoveDian00(priceStr string) string  {
	return strings.TrimRight(priceStr, ".00")
}
