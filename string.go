package gosupport

import "strings"

//gosupport.UrlDeal("nfangbian.com/fangan/index/?xyz=1#ab", "a=1&b=2")
func UrlDeal(reqUrl string, otherGetParam string) string {
	ret := ""
	suffix := ""
	if index := strings.Index(reqUrl, "#");index>=0 {
		ret = reqUrl[:index]
		suffix = reqUrl[index:]
	} else {
		ret = reqUrl
	}
	if strings.ContainsRune(reqUrl, '?') {
		ret += "&" + otherGetParam + suffix
	} else {
		ret += "?" + otherGetParam + suffix
	}
	return ret
}

