package setting

import (
	"golang-admin/internal/model"
	"gorm.io/gorm"
)

var tabs = []interface{}{
	model.User{},
	model.News{},
	model.Tag{},
}

// CreatTables 创建项目需要用到的表，如已存在则忽略
func CreatTables(db *gorm.DB) error {
	return db.AutoMigrate(tabs...)
}
