package gosupport

import (
	"fmt"
	"testing"
	"time"
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

func TestInitStruct4DefaultTag(t *testing.T) {
	config := &struct {
		B bool          `default:"true"` //true
		B2 bool         `default:"false"` //false
		S string        `default:"默认字符串哈哈" json:"s"`
		S2 string
		I int           `default:"10"`  //10
		I2 int8           `default:"7"` //7
		I3 int16           `default:"160"` //160
		I4 int32           `default:"320"`
		I5 int64           `default:"640"`
		F float64           `default:"9.8"`
		F2 float32           `default:"3.89"`
		T time.Duration `default:"1000000"` //1ms
		E int           `default:""`  //0
		E2 int           `default:"0"` //0
		U uint           `default:"11"` //11
		U2 uint8           `default:"22"` //22

	}{}
	InitStruct4DefaultTag(config)
	fmt.Println(fmt.Sprintf("%+v", config))
}

// go test -run="TestNewSequence"
func TestNewSequence(t *testing.T) {
	ret := NewSequence().SetPrefix("SO").SetStepSize(108).SetUid(10010).GetSeq()
	fmt.Println(ret)

	ret2 := NewSequence().SetPrefix("SO").SetStepSize(-1).SetUid(10010).GetSeq()
	fmt.Println(ret2)

	uidMod1 := SeqString(ret2).GetUidMod4Seq()
	ret3 := NewSequence().SetPrefix("SOL").SetStepSize(108).GetSeq(uidMod1)
	fmt.Println(ret3)

	var uidMod int64 = 789
	ret4 := NewSequence().SetPrefix("OPL").SetStepSize(108).GetSeq(uidMod)
	fmt.Println(ret4)

}
