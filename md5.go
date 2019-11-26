package gosupport

import (
	"crypto/md5"
	"fmt"
)

//md5加密
func Md5V1(str string) string {
	h := md5.New()  //返回hash.Hash接口对象
	h.Write([]byte(str))  //字符串转byte数组,并写入内容
	has := h.Sum([]byte(""))  //返回[]byte
	md5str := fmt.Sprintf("%x", has)
	return string(md5str)
}

