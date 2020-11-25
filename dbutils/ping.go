package dbutils

import "errors"
var (
	//表名不能为空
	ErrTableEmpty = errors.New("The table name cannot be empty")
	//插入内容不能为空
	ErrInsertEmpty = errors.New("The insert cannot be empty")
	//更新内容不能为空
	ErrUpdateEmpty = errors.New("The update cannot be empty")
)

