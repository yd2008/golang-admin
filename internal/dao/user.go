package dao

import (
	"golang-admin/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) CountUser(username string) (int64, error) {
	user := model.User{Username: username}
	return user.Count(d.engine)
}

func (d *Dao) CreateUser(username, password string, sex uint8) error {
	user := model.User{
		Username: username,
		Sex:      sex,
		Password: password,
	}

	return user.Create(d.engine)
}

func (d *Dao) QueryUser(query map[string]interface{}) (*model.User, error) {
	user := model.User{}
	return user.Query(d.engine, query)
}

func (d *Dao) GetUser(id uint) (*model.User, error) {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	return user.Get(d.engine)
}

func (d *Dao) DeleteUser(id uint) error {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	return user.Delete(d.engine)
}
