package model

import (
	"github.com/go-sql-driver/mysql"
	"golang-admin/pkg/errcode"
	"golang-admin/pkg/util"
	"gorm.io/gorm"
)

type User struct {
	Common
	Username      string `json:"username" gorm:"column:username;unique;type:varchar(100);comment:用户名"`
	Phone         string `json:"phone" gorm:"column:phone;unique;type:varchar(11);default:null;comment:电话"`
	Gender        uint8  `json:"gender" gorm:"column:gender;type:int(2);default:0;comment:性别"`
	Password      string `json:"-" gorm:"column:password;type:varchar(100);comment:密码;<-"`
	Avatar        string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像;"`
	IsWechatLogin uint8  `json:"is_wechat_login" gorm:"column:is_wechat_login;type:int(2);default:0;comment:是否是微信登录"`
	WechatId      string `json:"wechat_id" gorm:"column:wechat_id;unique;default:null;comment:微信openid"`
	Salt          string `json:"-" gorm:"column:salt;type:varchar(255);comment:加盐;<-"`
	Tags          string `json:"tags" gorm:"column:tag;type:varchar(255);default:'';comment:标签;"`
}

func (User) TableName() string {
	return "go_user"
}

func (u User) Count(db *gorm.DB) (int64, error) {
	var count int64
	if u.Username == "" {
		db.Where("username = ?", u.Username)
	}
	if err := db.Model(User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (u User) Create(db *gorm.DB) error {
	err := db.Create(&u).Error
	if sqlErr, ok := err.(*mysql.MySQLError); ok {
		err = errcode.RegisterUserFail.WithDetails(sqlErr.Error())
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
		if err.Error() == "record not found" {
			err = errcode.RegisterUserFail.WithDetails("未找到记录")
		}
		return nil, err
	}

	return &user, nil
}

func (u User) Query(db *gorm.DB, username, password string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
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
		db.Where("username = ?", u.Username)
	}
	if err = db.Offset(pageOffset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u User) WechatLogin(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("wechat_id = ?", u.WechatId).First(&user).Error
	if err != nil && err.Error() == "record not found" {
		if err = u.Create(db); err != nil {
			return nil, err
		}
		return &user, nil
	} else if err != nil {
		return nil, err
	}

	return &user, err
}
