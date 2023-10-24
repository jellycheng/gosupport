package gosupport

import (
	"strconv"
	"strings"
)

// -1小于，0相等，1大于
func CompareVersion(v1 string, v2 string) int {
	var res int
	v1 = strings.TrimLeft(v1, "Vv")
	v2 = strings.TrimLeft(v2, "Vv")
	v1Slice := strings.Split(v1, ".")
	v2Slice := strings.Split(v2, ".")
	v1Len := len(v1Slice)
	v2Len := len(v2Slice)
	verLen := v1Len
	if len(v1Slice) < len(v2Slice) {
		verLen = v2Len
	}
	for i := 0; i < verLen; i++ {
		var ver1Int, ver2Int int
		if i < v1Len {
			ver1Int, _ = strconv.Atoi(v1Slice[i])
		}
		if i < v2Len {
			ver2Int, _ = strconv.Atoi(v2Slice[i])
		}
		if ver1Int < ver2Int {
			res = -1
			break
		}
		if ver1Int > ver2Int {
			res = 1
			break
		}
	}
	return res
}
