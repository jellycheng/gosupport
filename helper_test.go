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
