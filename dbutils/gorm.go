package dbutils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

type MysqlGormInstance struct {
	mysql  map[string]*gorm.DB

}
//公有
func (mysqlInstance *MysqlGormInstance) GetMysql(dsnKey string) *gorm.DB {

	if d, ok := mysqlInstance.mysql[dsnKey]; ok {
		return d
	}
	return nil

}

//私有
func (mysqlInstance *MysqlGormInstance) registerMysql(dsn string, db *gorm.DB) *gorm.DB {
	mysqlInstance.mysql[dsn] = db
	return db
}


var mysqlGormObj = MysqlGormInstance{}
	//mysqlGormObj.mysql = make(map[string]*gorm.DB)

//实例化
func NewMysqlGorm(mysqlDsn MysqlDsn) *gorm.DB {
	if mysqlGormObj.mysql == nil{
		mysqlGormObj.mysql = make(map[string]*gorm.DB)
	}

	if d := mysqlGormObj.GetMysql(mysqlDsn.Key()); d != nil {
		return d
	}
	//实例化
	if d, err := gorm.Open("mysql", mysqlDsn.ToDsn()); err != nil {
		fmt.Println("connect mysql error: " + err.Error())
		return nil
	} else {//注册
		//registerScopes(d)
		return mysqlGormObj.registerMysql(mysqlDsn.Key(), d)
	}

}

