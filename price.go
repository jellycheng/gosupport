package gosupport

import (
	"fmt"
	"strings"
)

/**
	分转元
	price := 1909
	s1 := fen2yuan(int64(price), true)
 */
func fen2yuan(price int64, isTrimZero bool) string {
	tmpPrice01 := price/100
	tmpPrice02 := price - tmpPrice01 * 100
	formatPrice := ""
	if tmpPrice02<10 && tmpPrice02 >0 {
		formatPrice = fmt.Sprintf("%d.0%d", tmpPrice01, tmpPrice02)
	} else if tmpPrice02 == 0 {
		formatPrice = fmt.Sprintf("%d", tmpPrice01)
	} else {
		formatPrice = fmt.Sprintf("%d.%d", tmpPrice01, tmpPrice02)
	}

	if isTrimZero {
		formatPrice = strings.TrimRight(formatPrice, "0")
	}

	return formatPrice
}

