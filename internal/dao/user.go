package dao

import (
	"golang-admin/internal/model"
)

func (d *Dao) CountUser(username string) (int64, error) {
	user := model.User{Username: username}
	return user.Count(d.engine)
}

func (d *Dao) CreateUser(username, password string, gender uint8) error {
	user := model.User{
		Username: username,
		Gender:   gender,
		Password: password,
	}
	return user.Create(d.engine)
}

func (d *Dao) QueryUser(username, password string) (*model.User, error) {
	user := model.User{}
	return user.Query(d.engine, username, password)
}

func (d *Dao) GetUser(id uint) (*model.User, error) {
	var user = model.User{
		Common: model.Common{
			ID: id,
		},
	}
	return user.Get(d.engine)
}

func (d *Dao) DeleteUser(id uint) error {
	var user = model.User{
		Common: model.Common{
			ID: id,
		},
	}
	return user.Delete(d.engine)
}

