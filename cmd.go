package gosupport

import (
	"bytes"
	"os/exec"
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

//当前目录作为工作目录来执行命令
func ExecCmdBytes(cmdName string, args ...string) ([]byte, []byte, error) {
	return ExecCmdDirBytes("", cmdName, args...)
}

func ExecCmdDir(dir, cmdName string, args ...string) (string, string, error) {
	bufOut, bufErr, err := ExecCmdDirBytes(dir, cmdName, args...)
	return string(bufOut), string(bufErr), err
}

//当前目录作为工作目录来执行命令,调用示例： okContent,errContent, err := gosupport.ExecCmd("ls", "-l")
func ExecCmd(cmdName string, args ...string) (string, string, error) {
	return ExecCmdDir("", cmdName, args...)
}

// 切换到指定目录拉取最新代码
func GitPull(codeRootDir string) (string, error) {
	if okCon, failCon, err := ExecCmdDir(codeRootDir, "git", "pull"); err == nil {
		return okCon, nil
	} else {
		return failCon, err
	}
}

// 切换到指定目录checkout
func GitCheckout(codeRootDir string) (string, error) {
	if okCon, failCon, err := ExecCmdDir(codeRootDir, "git", "checkout", "--", "."); err == nil {
		return okCon, nil
	} else {
		return failCon, err
	}
}
