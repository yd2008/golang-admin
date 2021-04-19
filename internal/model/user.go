package model

import (
	"github.com/go-sql-driver/mysql"
	"golang-admin/pkg/errcode"
	"golang-admin/pkg/util"
	"gorm.io/gorm"
)

type User struct {
	Common
	Username string `json:"user_name" gorm:"column:user_name;unique;type:varchar(100);comment:用户名"`
	Sex      uint8  `json:"sex" gorm:"column:sex;type:int(2);default:0;comment:性别"`
	Password string `json:"-" gorm:"column:password;type:varchar(100);comment:密码;<-"`
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像;"`
	Salt     string `json:"-" gorm:"column:salt;type:varchar(255);comment:加盐;<-"`
}

func (User) TableName() string {
	return "go_user"
}

func (u User) Count(db *gorm.DB) (int64, error) {
	var count int64
	if u.Username == "" {
		db.Where("user_name = ?", u.Username)
	}
	if err := db.Model(User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (u User) Create(db *gorm.DB) error {
	err := db.Create(&u).Error
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		err = errcode.RegisterUserFail.WithDetails(sqlError.Error())
	}

	return err
}

func (u User) Delete(db *gorm.DB) error {
	return db.Delete(&u).Error
}

func (u User) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&u).Updates(values).Error
}

func (u User) Get(db *gorm.DB) (*User, error) {
	var user User
	if err := db.First(&user, u.ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Query(db *gorm.DB, username, password string) (*User, error) {
	var user User
	err := db.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = util.CompareWithPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if u.Username != "" {
		db.Where("user_name = ?", u.Username)
	}
	if err = db.Offset(pageOffset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
