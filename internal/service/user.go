package service

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/convert"
)

type RegisterUserReq struct {
	UserName string `json:"user_name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Sex      uint8  `json:"sex"`
}

type LoginUserReq struct {
	UserName string `json:"user_name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type GetUserReq struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

type DeleteUserReq struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

func (svc *Service) UserRegister(param *RegisterUserReq) error {
	return svc.dao.CreateUser(param.UserName, param.Password, param.Sex)
}

func (svc *Service) UserLogin(param *LoginUserReq) (*model.User, error) {
	query, err := convert.Struct2UnderlineMap(param)
	if err != nil {
		return nil, err
	}
	return svc.dao.QueryUser(query)
}

func (svc *Service) UserDelete(param *DeleteUserReq) error {
	return svc.dao.DeleteUser(param.ID)
}

func (svc *Service) UserGet(param *GetUserReq) (*model.User, error) {
	return svc.dao.GetUser(param.ID)
}