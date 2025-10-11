package gosupport

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 对pid文件的操作
type PIDFile struct {
	//pid文件
	pidFile string
}

// 创建pid文件,存在pid文件则写失败, pid=os.Getpid()
func (f PIDFile) CreatePidFile(pid int) error {
	path := filepath.Dir(f.pidFile)
	//判断目录是否存在
	if IsDir(path) == false {
		if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
			return err
		}
	}
	if IsFile(f.pidFile) == true {
		return errors.New(fmt.Sprintf("pid文件%s已经存在", f.pidFile))
	}
	if err := ioutil.WriteFile(f.pidFile, []byte(fmt.Sprintf("%d", pid)), 0644); err != nil {
		return err
	}
	return nil
}

// 移除pid文件
func (f PIDFile) Remove() error {
	return os.Remove(f.pidFile)
}

func (f PIDFile) GetPid() (int, error) {
	if pid, err := FileGetContents(f.pidFile); err == nil {
		pid = strings.TrimSpace(pid)
		return strconv.Atoi(pid)
	} else {
		return 0, err
	}

}

/*
调用示例：
pidFile := "/data/www/go-community-service/go-community-service.pid"
pidFileObj := gosupport.NewPIDFile(pidFile)
err := pidFileObj.CreatePidFile(os.Getpid())

	if err !=nil {
		fmt.Println(err.Error())
	}

pid2,_ := pidFileObj.GetPid()
fmt.Println("pidfile=", pidFile, "pid=", pid2)
pidFileObj.Remove()
*/
func NewPIDFile(pidfile string) *PIDFile {
	pidfile = strings.TrimSpace(pidfile)
	return &PIDFile{pidFile: pidfile}
}
