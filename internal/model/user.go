package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"user_name" gorm:"column:username;type:varchar(100);comment:用户名"`
	Sex      uint8  `json:"sex" gorm:"column:sex;type:int(2);comment:性别"`
	Password string `json:"password" gorm:"column:password;type:varchar(100);comment:密码;<-"`
	Salt     string `json:"-" gorm:"column:salt;type:varchar(255);comment:加盐;<-"`
}

func (User) TableName() string {
	return "user"
}
