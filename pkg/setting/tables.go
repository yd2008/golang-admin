package setting

import (
	"golang-admin/internal/model"
	"gorm.io/gorm"
)

var tabs = []interface{}{model.User{}}

// 创建项目需要用到的表，如已存在则忽略
func CreatTables(db *gorm.DB) error {
	err := db.AutoMigrate(tabs...)
	if err != nil {
		return err
	}
	return nil
}
