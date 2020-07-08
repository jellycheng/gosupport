package dbutils

import (
	"fmt"
	"testing"
)

type SystemModel struct {
	ID        uint64 `gorm:"primary_key;Column:id"`
	IsDelete  uint8 `gorm:"Column:is_delete;DEFAULT:0"`
	CreateTime uint64 `gorm:"Column:create_time"`
	UpdateTime uint64 `gorm:"Column:update_time"`
	DeleteTime uint64 `gorm:"Column:delete_time"`

	SystemCode string `gorm:"Column:system_code"`
	SystemName string `gorm:"Column:system_name"`
	Appid string `gorm:"Column:app_id"`
	Secret string `gorm:"Column:secret"`
}

func (SystemModel)TableName() string {
	return "t_system"
}

func TestNewMysqlGorm222(t *testing.T) {
	//数据库配置
	mysqlDsnObj := NewMysqlDsn(map[string]interface{}{
		"host":"10.30.60.122",
		"username":"devreadonly",
		"password":"ug<H7+hsSmw1",
		"port":3306,
		"dbname":"towngas_daojia_common",
	})
	//打印dsn串
	fmt.Println(mysqlDsnObj.ToDsn())
	//根据db配置获取*gorm.DB对象
	gormDb := NewMysqlGorm(*mysqlDsnObj)
	//gorm设置 todo，如 log、debug、连接配置等

	//执行查询sql: SELECT * FROM `t_system`  WHERE (system_name='xxx后台管理')
	var systemModel SystemModel
	gormDb.Debug().Where("system_name=?", "xxx后台管理").Find(&systemModel)
	fmt.Println(systemModel)

}
