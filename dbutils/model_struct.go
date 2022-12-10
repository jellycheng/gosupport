package dbutils

type GormModel struct {
	ID         int64 `gorm:"primary_key;Column:id"`
	IsDelete   int8  `gorm:"Column:is_delete;DEFAULT:0"`
	CreateTime int64 `gorm:"Column:create_time"`
	UpdateTime int64 `gorm:"Column:update_time"`
	DeleteTime int64 `gorm:"Column:delete_time"`
}
