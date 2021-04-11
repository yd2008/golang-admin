package setting

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(setting *DatabaseSettingS) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", setting.UserName, setting.Password, setting.Host, setting.DBName, setting.Charset, setting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
