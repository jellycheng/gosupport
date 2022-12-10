package dbutils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	//表名不能为空
	ErrTableEmpty = errors.New("The table name cannot be empty")
	//插入内容不能为空
	ErrInsertEmpty = errors.New("The insert cannot be empty")
	//更新内容不能为空
	ErrUpdateEmpty = errors.New("The update cannot be empty")
)

//in, v := PinWhereIn("username", "in", []interface{}{11, 22,33}) 返回 username IN (?,?,?) 、[11 22 33]
func PinConditionIn(field string, condition string, values []interface{}) (string, []interface{}) {
	var buf strings.Builder
	wps := WenHaoPlaceholders(len(values))
	buf.WriteString(fmt.Sprintf(" %s %s (%s)", field, condition, wps))
	newWhere := buf.String()
	return newWhere, values
}

//in, v := PinWhereNotIn("username", []interface{}{11, 22,33}) 返回 username NOT IN (?,?,?) 、[11 22 33]
func PinWhereNotIn(field string, values []interface{}) (string, []interface{}) {
	var buf strings.Builder
	wps := WenHaoPlaceholders(len(values))
	buf.WriteString(fmt.Sprintf(" %s NOT IN (%s)", field, wps))
	newWhere := buf.String()
	return newWhere, values
}
