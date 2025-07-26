package gosupport

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

// 写内容
func FilePutContents(filename string, content string, perm ...os.FileMode) (int, error) {
	var fileMode os.FileMode
	if len(perm) == 0 {
		fileMode = 0666
	} else {
		fileMode = perm[0]
	}
	if f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileMode); err == nil {
		defer f.Close()
		n, err2 := io.WriteString(f, content)
		return n, err2
	} else {
		return 0, err
	}
}

// 判断文件/文件夹是否存在，true存在，false不存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) { //只是没有权限
			return true
		}
		return false
	}
	return true
}

// 判断文件是否存在
func IsFile(filename string) bool {
	s, err := os.Stat(filename)
	if err != nil { //不是文件也不是目录
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

// 获取文件大小
func FileSize(filename string) (int64, error) {
	f, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

func FileExistAndSize(filename string) (int64, bool, error) {
	f, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return f.Size(), true, nil
}

// 获取文件修改时间
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

func IsDirWriteable(dir string) error {
	f := filepath.Join(dir, ".touch")
	if err := ioutil.WriteFile(f, []byte(""), PrivateFileMode); err != nil {
		return err
	}
	return os.Remove(f)
}

/*
调用示例：

	data := make(map[string]interface{})
	if err :=LoadJson("./cjs.json", &data);err!=nil{
		fmt.Println("解析失败", err.Error())
	} else {
		fmt.Println(data)
	}
*/
func LoadJson(jsonFile string, t interface{}) error {
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	return nil
}

// 获取文件扩展名
func FileExt(f string) string {
	ext := filepath.Ext(f)
	if ext == "" {
		return ext
	}
	return ext[1:]
}

func FileBaseName(f string) string {
	n := filepath.Base(f)
	return n
}

// 创建上级目录,存在或者创建成功返回true、创建失败返回false
func CreateSuperiorDir(path string) bool {
	if path == "" || StrInSlice(path, []string{".", "..", "./", "../", "/"}) {
		return false
	}
	spath := filepath.Dir(path)
	if spath == "" {
		return false
	}
	if err := os.MkdirAll(spath, 0755); err == nil {
		return true
	} else {
		return false
	}
}

// 遍历目录
func WalkDirRecursive(dir string, isIgnoreHide bool) ([]string, []string, error) {
	fileList := []string{}
	dirList := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fileList, dirList, err
	}

	for _, f := range files {
		if isIgnoreHide && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if f.IsDir() {
			// 当前目录
			dirList = append(dirList, filepath.Join(dir, f.Name()))

			// 递归调用处理子目录
			subDirPath := filepath.Join(dir, f.Name())
			subFileList, subDirList, err := WalkDirRecursive(subDirPath, isIgnoreHide)
			if err != nil {
				return fileList, dirList, err
			}
			// 将子目录中的文件和目录添加到当前列表中
			fileList = append(fileList, subFileList...)
			dirList = append(dirList, subDirList...)
		} else {
			fileList = append(fileList, filepath.Join(dir, f.Name()))
		}
	}
	return fileList, dirList, nil
}
