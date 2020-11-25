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
	sqlObj := NewSQLBuilderUpdate().SetTable("bbs_user").SetUpdateData([]string{"qq"}, 123, 444).AndWhere("uid","=",5)
	sql, _ := sqlObj.GetSQL()
	fmt.Println(sql)

	sqlObj2 := NewSQLBuilderUpdate().SetTable("bbs_user").SetUpdateData([]string{"qq","username"}, 123, 444).AndWhere("uid","=",5)
	sqlObj2.SetLimit("3")
	sqlObj2.SetOrderBy("uid asc")
	sql2, _ := sqlObj2.GetSQL()
	fmt.Println(sql2, sqlObj2.GetParamValues(), sqlObj2.GetSetParamValues(),sqlObj2.GetWhereParamValues())

}

//go test -run="TestDeleteSql"
func TestDeleteSql(t *testing.T) {
	sqlObj1 := NewSQLBuilderDelete().SetTable("bbs_user")
	sql1, _ := sqlObj1.GetSQL()
	fmt.Println(sql1)

	sqlObj2 := NewSQLBuilderDelete().SetTable("bbs_user").SetLimit("2").SetOrderBy("uid asc")
	sql2, _ := sqlObj2.GetSQL()
	fmt.Println(sql2)

	sqlObj3 := NewSQLBuilderDelete().SetTable("bbs_user").AndWhere("uid", "=", 9)
	sqlObj3.AndWhere("username", "=", "admin")
	sqlObj3.SetLimit("3").SetOrderBy("uid desc")
	sqlObj3.OrWhere("qq", "=", 123)
	sql3, _ := sqlObj3.GetSQL()
	fmt.Println(sql3, sqlObj3.GetWhereParamValues())



}

