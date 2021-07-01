package xlog

import (
	"bytes"
	"log"
	"os"
	"testing"
	"fmt"
)

// go test --run="TestInfo"
func TestInfo(t *testing.T) {
	Debug("debug级别日志内容")
	Info("info级别日志内容")
	Warn("Warn级别日志内容")
	Error("error级别日志内容")

	infoLogObj := New(os.Stdout, "[user-service] ", log.Lmicroseconds | log.LstdFlags | log.Llongfile, InfoLevel)
	infoLogObj.Trace("trace信息")
	infoLogObj.Debug("debug信息")
	infoLogObj.Info("info信息")
	infoLogObj.Warn("warn信息")
	infoLogObj.Error("error信息")
	//infoLogObj.Fatal("fatal信息，打印完，退出")
	//infoLogObj.Panic("panic信息")
	infoLogObj.Printf("yes：%s", "go原生方法也可以调用")
	goLogObj := infoLogObj.GetGoLogger()
	goLogObj.Println("使用原生log日志打印")

	b := bytes.Buffer{}
	loggerObj := New(&b, "order-service ", log.LstdFlags | log.Llongfile, TraceLevel)
	loggerObj.Info("hello world")
	fmt.Printf("日志信息: %s \n", b.String())

	levelUint8,_ := ParseLevel("error")
	fmt.Println(levelUint8) //4

	levelUint82,_ := ParseLevel("no")
	fmt.Println(levelUint82) //0

	fmt.Println(WarnLevel.ToString()) // warn <nil>

}
