package service

import (
	"golang-admin/internal/dao"
	"golang-admin/pkg/util"
)

type RegisterUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" binding:"required,min=11,max=11"`
	Gender   uint8  `json:"gender" binding:"oneof=0 1"`
}

type UpdateUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" binding:"required,min=11,max=11"`
	Gender   uint8  `json:"gender" binding:"oneof=0 1"`
}

type LoginUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type DeleteUserReq struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

func (svc *Service) UserRegister(param *RegisterUserBody) error {
	encryptPwd, err := util.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}
	return svc.dao.CreateUser(param.Username, encryptPwd, param.Phone, param.Avatar, param.Gender)
}

func (svc *Service) UserUpdate(id uint,param *UpdateUserBody) error {
	return svc.dao.UpdataUser(id, param.Username, param.Phone, param.Avatar, param.Gender)
}

func (svc *Service) UserLogin(param *LoginUserBody) (*dao.User, error) {
	return svc.dao.QueryUser(param.Username, param.Password)
}

func (svc *Service) UserDelete(param *DeleteUserReq) error {
	return svc.dao.DeleteUser(param.ID)
}

func (svc *Service) UserGet(id uint) (*dao.User, error) {
	return svc.dao.GetUser(id)
}
