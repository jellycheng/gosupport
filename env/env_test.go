package env

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"os"
	"strings"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	var envPath string = "/Users/jelly/test/mygoenv/"

	err := LoadEnv(envPath + ".env", envPath + ".env.local")
	if err!=nil {
		fmt.Println(err.Error())
	}

	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		tmp := strings.Split(rawEnvLine, "=")
		println("key:", tmp[0], "val:",tmp[1])
	}
	fmt.Println("=================================")

	//加载env文件覆盖已存在相同的环境变量值
	Overload(envPath + ".env.local")
	rawEnv2 := os.Environ()
	for _, rawEnvLine := range rawEnv2 {
		tmp := strings.Split(rawEnvLine, "=")
		println(tmp[0], " = ",tmp[1])
	}

	fmt.Println("=================================")
	println("app_id=", Get("app_id", "no appid "))
}



func TestLoadEnv2DataManage(t *testing.T) {
	var envPath string = "/Users/jelly/test/mygoenv/"

	err := LoadEnv2DataManage(envPath + ".env", envPath + ".env.local")
	if err!=nil {
		fmt.Println(err.Error())
	}
	globalenv := gosupport.NewGlobalEnvSingleton()
	data := globalenv.GetData()
	for k,v := range data {
		fmt.Println("key:", k, " val:", v)
	}

	fmt.Println("=================================")
	fmt.Println(globalenv.Data["USER_1_DB_READ_HOST"])

}

