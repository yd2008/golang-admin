package service

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/util"
)

type RegisterUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Phone    string `json:"phone" binding:"required,min=11,max=11"`
	Gender   uint8  `json:"gender"`
}

type LoginUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type GetUserReq struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

type DeleteUserReq struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

func (svc *Service) UserRegister(param *RegisterUserBody) error {
	encryptPwd, err := util.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}
	return svc.dao.CreateUser(param.Username, encryptPwd, param.Phone, param.Gender)
}

func (svc *Service) UserLogin(param *LoginUserBody) (*model.User, error) {
	return svc.dao.QueryUser(param.Username, param.Password)
}

func (svc *Service) UserDelete(param *DeleteUserReq) error {
	return svc.dao.DeleteUser(param.ID)
}

func (svc *Service) UserGet(param *GetUserReq) (*model.User, error) {
	return svc.dao.GetUser(param.ID)
}
