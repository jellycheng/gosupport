package gosupport

import "fmt"

type WrapError struct {
	Error error
}

// AddError 追加db错误信息对象
func (me *WrapError)AddError(e error) error {
	if me.Error == nil{
		me.Error = e
	} else {
		me.Error = fmt.Errorf("%v; %w", me.Error, e)
	}
	return me.Error
}

func NewWrapError() WrapError {
	return WrapError{}
}
