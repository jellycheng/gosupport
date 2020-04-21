package dbutils

type GormModel struct {
	ID         uint64 `gorm:"primary_key;Column:id"`
	IsDelete   uint8 `gorm:"Column:is_delete;DEFAULT:0"`
	CreateTime uint64 `gorm:"Column:create_time"`
	UpdateTime uint64 `gorm:"Column:update_time"`
	DeleteTime uint64 `gorm:"Column:delete_time"`

}


