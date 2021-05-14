package dao

import (
	"golang-admin/internal/model"
)

type User struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Phone         string `json:"phone"`
	Gender        uint8  `json:"gender"`
	Avatar        string `json:"avatar"`
	IsWechatLogin uint8  `json:"is_wechat_login"`
	WechatId      string `json:"wechat_id"`
}

func (d *Dao) CountUser(username string) (int64, error) {
	user := model.User{Username: username}
	return user.Count(d.engine)
}

func (d *Dao) CreateUser(username, password, phone, avatar string, gender uint8) error {
	user := model.User{
		Username: username,
		Gender:   gender,
		Password: password,
		Avatar:   avatar,
		Phone:    phone,
	}
	return user.Create(d.engine)
}

func (d *Dao) UpdataUser(id uint, username, phone, avatar string, gender uint8) error {
	var user = model.User{
		Common: model.Common{
			ID: id,
		},
	}
	var updataInfo = make(map[string]interface{})
	updataInfo["username"] = username
	updataInfo["phone"] = phone
	updataInfo["avatar"] = avatar
	updataInfo["gender"] = gender
	return user.Update(d.engine, updataInfo)
}

func (d *Dao) QueryUser(username, password string) (*User, error) {
	_user := model.User{}
	user, err := _user.Query(d.engine, username, password)
	if err != nil {
		return nil, err
	}

	return createUserDao(user), nil
}

func (d *Dao) GetUser(id uint) (*User, error) {
	var _user = model.User{
		Common: model.Common{
			ID: id,
		},
	}
	user, err := _user.Get(d.engine)
	if err != nil {
		return nil, err
	}

	return createUserDao(user), nil
}

func (d *Dao) DeleteUser(id uint) error {
	var user = model.User{
		Common: model.Common{
			ID: id,
		},
	}
	return user.Delete(d.engine)
}

func createUserDao(user *model.User) *User {
	var userDao = User{
		ID:            user.Common.ID,
		Username:      user.Username,
		Phone:         user.Phone,
		Gender:        user.Gender,
		Avatar:        user.Avatar,
		IsWechatLogin: user.IsWechatLogin,
		WechatId:      user.WechatId,
	}
	return &userDao
}
