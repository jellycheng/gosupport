package gosupport

import (
	"io"
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

//写内容
func FilePutContents(filename string, content string, perm ...os.FileMode) (int, error) {
	var fileMode os.FileMode
	if (len(perm) == 0) {
		fileMode = 0666
	} else {
		fileMode = perm[0]
	}
	if f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileMode); err == nil{
		n, err2 := io.WriteString(f, content)
		return n, err2
	} else {
		return 0, err
	}
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

//获取文件大小
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

//获取文件修改时间
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}


