package gosupport

import (
	"io/ioutil"
	"os"
)

// 读取文件内容
func FileGetContents(filename string) (string, error) {
	//io/ioutil包：func ReadFile(filename string) ([]byte, error) 读取整个文件内容
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 判断文件/文件夹是否存在，true存在，false不存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) { //只是没有权限
			return true
		}
		return false
	}
	return true
}

// 判断文件是否存在
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {//不是文件也不是目录
		return false
	}
	return !s.IsDir()
}

// 判断是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
