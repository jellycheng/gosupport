package gosupport

import (
	"fmt"
	"testing"
)

func TestIntSum(t *testing.T) {

	total := IntSum(1, 3, 5)
	fmt.Println(total) //9

	//切片单元值求和
	nums := []int{1, 2, 3, 4}
	fmt.Println(IntSum(nums...)) //10


}

func TestToJson(t *testing.T) {
	type UserInfo struct {
		Userid  int         `json:"userid"`
		Nickname   string      `json:"nickname"`
		Age  int         `json:"age"`
	}

	//{"userid":123,"nickname":"张三","age":28}
	fmt.Println(ToJson(&UserInfo{Userid:123, Age:28, Nickname:"张三"}))

}
