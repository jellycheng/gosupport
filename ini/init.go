package ini

import (
	"bytes"
	"regexp"
	"runtime"
)

const (
	//默认组名
	DefaultGroupName = "default"
)

var (
	//换行符
	LineBreak = "\n"

	//abc = gopkg.in/%(NAME)s.%(VERSION)s
	VarPattern = regexp.MustCompile(`%\(([^)]+)\)s`)
)

func init() {

	if runtime.GOOS == "windows" {
		LineBreak = "\r\n"
	}

}

//获取注释内容
func GetComment(in []byte) ([]byte, bool) {
	i := bytes.IndexAny(in, "#;")
	if i == -1 {
		return nil, false
	}
	return in[i:], true
}

//去除注释内容
func GetCleanComment(in []byte) []byte {
	i := bytes.IndexAny(in, "#;")
	if i == -1 {
		return in
	}
	return in[:i]
}
