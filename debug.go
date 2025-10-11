package gosupport

import (
	"fmt"
	"reflect"
	"strings"
)

// 调用示例：gosupport.DebugPrintReflect(time.Hour)
func DebugPrintReflect(x interface{}) {
	v := reflect.ValueOf(x) //反射，返回reflect.Value结构体对象
	t := v.Type()           //返回reflect.Type接口对象
	fmt.Printf("类型为 %s,类型支持的方法如下：\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type() //具体某个方法转成reflect.Type接口对象
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
