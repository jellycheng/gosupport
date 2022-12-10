package gosupport

import (
	"fmt"
	"regexp"
)

// ToRed 返回红色样式，用于在终端打印红色字，调用示例：fmt.Println(gosupport.ToRed("hello world"))
func ToRed(str string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", str)
}

func ToGreen(str string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", str)
}

func ToYellow(str string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", str)
}

func ToColor(code string, msg string) string {
	cTpl := "\033[%sm%s\033[0m"
	return fmt.Sprintf(cTpl, code, msg)
}

// ClearColorCode 提取颜色样式中的文本内容，示例：fmt.Println(ClearColorCode(ToYellow("goods")))
func ClearColorCode(str string) string {
	colorExpr := `\033\[[\d;?]+m`
	codeRegex := regexp.MustCompile(colorExpr)
	return codeRegex.ReplaceAllString(str, "")
}
