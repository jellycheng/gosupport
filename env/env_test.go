package env

import (
	"os"
	"strings"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	var envPath string = "/Users/jelly/test/mygoenv/"

	LoadEnv(envPath + ".env", envPath + ".env.local")
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		tmp := strings.Split(rawEnvLine, "=")
		println("key:", tmp[0], "val:",tmp[1])
	}

	//加载env文件覆盖已存在的环境变量值
	Overload(envPath + ".env.local")
	rawEnv2 := os.Environ()
	for _, rawEnvLine := range rawEnv2 {
		tmp := strings.Split(rawEnvLine, "=")
		println(tmp[0], " = ",tmp[1])
	}

	println("app_id=", Get("app_id", "no appid "))
}

