package utils

import "fmt"

//  错误结构体，并实现了error接口
//  myErr := utils.NewMyError()
//  myErr.ErrCode = 100
//  myErr.ErrMsg = "参数错误"
//  fmt.Println(myErr.Error())
//  fmt.Println(myErr.ErrCode)
//  fmt.Println(myErr.ErrMsg)
//  fmt.Print(myErr.ToString())
type MyError struct {
	ErrCode  int `json:"err_code"`  //错误代号
	ErrMsg   string `json:"err_msg"`  //错误信息
	ExtData  interface{}  `json:"ext_data"` // 错误扩展数据
}

func NewMyError() MyError {
	return MyError{}
}

func (m MyError) Error() string {
	return m.ErrMsg
}

func (m MyError) ToString() string {
	return fmt.Sprintf("err_code:%d, err_msg: %s ext_data: %v \n", m.ErrCode, m.ErrMsg, m.ExtData)
}
