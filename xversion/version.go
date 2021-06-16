package xversion

import "fmt"

var (
	// 服务名
	ServiceName = ""
	// 项目版本信息
	Version = ""
	// Go版本信息
	GoVersion = ""
	// git提交commmit id
	GitCommit = ""
	// 项目构建时间
	BuildTime = ""
)

func GetInfo() string {
	ret := fmt.Sprintf(`Service Name: %s
Version: %s
Go Version: %s
Git Commit: %s
Build Time: %s
`, ServiceName,Version,GoVersion,GitCommit,BuildTime)
	return ret
}

