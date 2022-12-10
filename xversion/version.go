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
	// 仓库分支名
	BranchName = ""
	// 项目构建时间
	BuildTime = ""
)

func GetInfo() string {
	ret := fmt.Sprintf(`Service Name: %s
Version: %s
Go Version: %s
Git Commit: %s
Branch Name: %s
Build Time: %s
`, ServiceName, Version, GoVersion, GitCommit, BranchName, BuildTime)
	return ret
}
