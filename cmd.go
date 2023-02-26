package gosupport

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//go执行命令

func ExecCmdDirBytes(dir, cmdName string, args ...string) ([]byte, []byte, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir //工作目录，如果为空则为当前目录
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	err := cmd.Run()
	return bufOut.Bytes(), bufErr.Bytes(), err
}

// ExecCmdBytes 当前目录作为工作目录来执行命令
func ExecCmdBytes(cmdName string, args ...string) ([]byte, []byte, error) {
	return ExecCmdDirBytes("", cmdName, args...)
}

func ExecCmdDir(dir, cmdName string, args ...string) (string, string, error) {
	bufOut, bufErr, err := ExecCmdDirBytes(dir, cmdName, args...)
	return string(bufOut), string(bufErr), err
}

// ExecCmd 当前目录作为工作目录来执行命令,调用示例： okContent,errContent, err := gosupport.ExecCmd("ls", "-l")
func ExecCmd(cmdName string, args ...string) (string, string, error) {
	return ExecCmdDir("", cmdName, args...)
}

// GitPull 切换到指定目录拉取最新代码
func GitPull(codeRootDir string) (string, error) {
	if okCon, failCon, err := ExecCmdDir(codeRootDir, "git", "pull"); err == nil {
		return okCon, nil
	} else {
		return failCon, err
	}
}

// GitCheckout 切换到指定目录checkout
func GitCheckout(codeRootDir string) (string, error) {
	if okCon, failCon, err := ExecCmdDir(codeRootDir, "git", "checkout", "--", "."); err == nil {
		return okCon, nil
	} else {
		return failCon, err
	}
}

// GetExeSuffix 获取系统可执行文件扩展名
func GetExeSuffix() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

// LookPath 通过可执行文件名返回命令绝对路径, gosupport.LookPath("git")
func LookPath(binFile string) (string, error) {
	suffix := GetExeSuffix()
	if len(suffix) > 0 && !strings.HasSuffix(binFile, suffix) {
		binFile = binFile + suffix
	}
	bin, err := exec.LookPath(binFile)
	if err != nil {
		return "", err
	}
	return bin, nil
}

// GitCloneIntoPath 克隆git代码到指定目录下
func GitCloneIntoPath(url, branch, codeHome string) error {
	ext := filepath.Ext(url)                            // 文件扩张名
	repo := strings.TrimSuffix(filepath.Base(url), ext) // 仓库名
	codeDir := filepath.Join(codeHome, repo)
	if FileExists(codeDir) {
		_ = os.RemoveAll(codeDir)
	}
	binFile, err := LookPath("git")
	if err != nil {
		return err
	}
	args := []string{"clone"}
	if len(branch) > 0 {
		args = append(args, "-b", branch)
	}
	args = append(args, url, codeDir)
	cmd := exec.Command(binFile, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}
