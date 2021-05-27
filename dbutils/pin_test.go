package dbutils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

//go test -run="TestInsertSql"
func TestInsertSql(t *testing.T) {
	sqlObj := NewSQLBuilderInsert().SetTable("bbs_user").SetInsertData([]string{"`email`", "username","qq"}, "tom899@qq.com", "admin899", 12345)
	sql,_ := sqlObj.GetSql()
	fmt.Println(sql)
	fmt.Println(sqlObj.GetParamValues())

	if conn,err := DbConnect(GetDsn(map[string]interface{}{"host":"localhost","user_name":"root","password":"88888888","dbname":"xiuno4"}));err==nil{
		insertId, _ := InsertSql(conn, sql, sqlObj.GetParamValues()...)
		fmt.Println(insertId)
	} else {
		fmt.Println(err.Error())
	}


	sqlObj2 := NewSQLBuilderInsert().SetTable("bbs_user").SetInsertData([]string{"email"}, "tom8@qq.com", "admin1238", 12345)
	sql2, err:= sqlObj2.GetSql()
	fmt.Println(sql2, err,sqlObj2.GetParamValues())

}

//go test -run="TestUpdateSql"
func TestUpdateSql(t *testing.T) {
	sqlObj := NewSQLBuilderUpdate().SetTable("bbs_user").SetUpdateData([]string{"qq"}, 123, 444).Where("uid","=",5)
	sql, _ := sqlObj.GetSQL()
	fmt.Println(sql)

	sqlObj2 := NewSQLBuilderUpdate().SetTable("bbs_user").SetUpdateData([]string{"qq","username"}, 123, 444).Where("uid","=",5)
	sqlObj2.SetLimit("3")
	sqlObj2.OrderBy("uid asc")
	sqlObj2.Where("xyz", "!=", 6)
	sqlObj2.OrWhere("orxyz", "!=", 7)
	sqlObj2.WhereRaw("`rawtitle` = ?", "hello")
	sql2, _ := sqlObj2.GetSQL()
	fmt.Println(sql2, sqlObj2.GetParamValues(), sqlObj2.GetSetParamValues(),sqlObj2.GetWhereParamValues())

}

//go test -run="TestDeleteSql"
func TestDeleteSql(t *testing.T) {
	sqlObj1 := NewSQLBuilderDelete().SetTable("bbs_user")
	sql1, _ := sqlObj1.GetSQL()
	fmt.Println(sql1)

	sqlObj2 := NewSQLBuilderDelete().SetTable("bbs_user").SetLimit("2").OrderBy("uid asc")
	sql2, _ := sqlObj2.GetSQL()
	fmt.Println(sql2)

	sqlObj3 := NewSQLBuilderDelete().SetTable("bbs_user").Where(WrapField("uid"), "=", 9)
	sqlObj3.Where("username", "=", "admin")
	sqlObj3.SetLimit("3").OrderBy("uid desc")
	sqlObj3.OrWhere("qq", "=", 123)
	sqlObj3.WhereIn("abc", []interface{}{1, 3, 5}...)
	sqlObj3.WhereNotIn("abcnot", []interface{}{1, 3, 5}...)
	sqlObj3.OrWhereIn("orabc", []interface{}{1, 3, 5}...)
	sqlObj3.OrWhereNotIn("ornotabc", 4,5,6,7)
	sqlObj3.WhereRaw("`rawtitle` = ?", "hello")
	sqlObj3.OrWhereRaw("(`OrWhereRawage` = ? OR `OrWhereRawage` = ?) AND `OrWhereRawclass` = ?", 22, 25, "2-3")
	sql3, _ := sqlObj3.GetSQL()
	fmt.Println(sql3, sqlObj3.GetWhereParamValues())



}

//go test -run="TestSelectSql"
func TestSelectSql(t *testing.T) {
	sqlObj1 := NewSQLBuilderSelect().SetTable("t_user")
	sqlObj1.Select("`age`", "COUNT(age)")
	sqlObj1.GroupBy("`age`", "`id`")
	sqlObj1.OrderBy("create_time desc,id asc")
	sqlObj1.Where(WrapField("uid"), "=", 19)

	if sql,err := sqlObj1.GetSQL();err == nil {
		fmt.Println(sql, sqlObj1.GetWhereParamValues())
	} else {
		fmt.Println(err.Error())
	}

}

// go test -run="TestJoinInt2Str"
func TestJoinInt2Str(t *testing.T) {
	s := JoinInt2Str([]int{123456789, 8,0,1}, ",")
	fmt.Println(s)
	s2 := JoinInt642Str([]int64{66, -1,0,1}, ",")
	fmt.Println(s2)
}

