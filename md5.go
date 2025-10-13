package gosupport

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// md5加密
func Md5(str string) string {
	return Md5V2(str)
}

// md5加密
func Md5V1(str string) string {
	h := md5.New()           //返回hash.Hash接口对象
	h.Write([]byte(str))     //字符串转byte数组,并写入内容
	has := h.Sum([]byte("")) //返回[]byte
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 使用md5加密方式2
func Md5V2(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5V3(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func Md5V4(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}

// 调用示例： gosupport.Md5V5([]byte{'h','e', 'l','l','o',' ', 'w','o','r','l','d'})
func Md5V5(bt []byte) string {
	h := md5.New()
	h.Write(bt)
	return hex.EncodeToString(h.Sum(nil))
}

// 获取文件的md5值
func Md5File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 获取文件的MD5值
func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()
	hash := md5.New()
	// 将文件内容写入哈希计算器
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("计算哈希失败: %v", err)
	}
	// 计算哈希值并转换为16进制字符串
	md5Bytes := hash.Sum(nil)
	md5Str := hex.EncodeToString(md5Bytes)

	return md5Str, nil
}
