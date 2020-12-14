package gosupport

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

//md5加密
func Md5V1(str string) string {
	h := md5.New()  //返回hash.Hash接口对象
	h.Write([]byte(str))  //字符串转byte数组,并写入内容
	has := h.Sum([]byte(""))  //返回[]byte
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5V2(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func Md5V4(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}
