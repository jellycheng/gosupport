package utils

import "strings"

// 用户ID转唯一邀请码
type GetInviteCode struct {
	obj    MyintCompress
	posVal string
}

func NewGetInviteCode() GetInviteCode {
	str := "abcdefghkmnpqrstwxy345678"
	compObj := MyintCompress{
		CharStr: str,
	}
	return GetInviteCode{
		obj:    compObj,
		posVal: "9",
	}
}

func (m GetInviteCode) EnCode(userid int64) string {
	ret := ""
	if userid <= 0 {
		return ret
	}
	ret = m.obj.Compress(userid)
	curLen := len(ret)
	if curLen < 4 && m.posVal != "" {
		ret = strings.Repeat(m.posVal, 4-curLen) + ret
	}
	return ret
}

func (m GetInviteCode) DeCode(code string) int64 {
	code = strings.TrimLeft(code, m.posVal)
	return m.obj.UnCompress(code)
}
