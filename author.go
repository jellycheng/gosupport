package gosupport

import "fmt"

//作者格式结构体
type Author struct {
	Name  string //作者
	Email string //作者邮箱
}

//返回作者格式
func (a Author) String() string {
	e := ""
	if a.Email != "" {
		e = " <" + a.Email + ">"
	}
	return fmt.Sprintf("%v%v", a.Name, e)
}

/**
调用示例：
var Authors []gosupport.Author //存批量作者
Authors = append(Authors, gosupport.Author{Name: "张三", Email: "admin@xxx.com"})
Authors = append(Authors, gosupport.Author{Name: "李四", Email: "lisi@xxx.com"})
for _,v:=range Authors{
	fmt.Println(v.String())
}
 */
