package dbutils

import (
	"fmt"
	"testing"
)

// go test -run=TestShowCreateTable
func TestShowCreateTable(t *testing.T) {
	dsn := GetDsn(map[string]interface{}{"host":"localhost","port":"3306", "dbname":"db_user_1","username":"root","password":"88888888"})
	con, err := GetDbConnect("user", dsn)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	tbl, err := ShowCreateTable(con, "t_user_token_1")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tbl)

}

// go test -run=TestGetMysqlVersion
func TestGetMysqlVersion(t *testing.T) {
	dsn := GetDsn(map[string]interface{}{"host":"localhost","port":"3306", "dbname":"db_user_1","username":"root","password":"88888888"})
	con, err := GetDbConnect("user", dsn)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(GetMysqlVersion(con))
}

