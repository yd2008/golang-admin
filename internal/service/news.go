package service

import (
	"golang-admin/internal/dao"
	"golang-admin/pkg/app"
)

type CreateNewsBody struct {
	Title string `json:"title" binging:"required,min=3,max=100"`
	Content string `json:"content" binging:"required,min=3"`
	Tags string `json:"tags"`
}

type CountNewsReq struct {
	Title string `json:"title" binging:"max=100"`
}

type ListNewsReq struct {
	Title string `form:"title"`
}

func (svc *Service) NewsCreate(param *CreateNewsBody) error {
	return svc.dao.CreateNews(param.Title, param.Content, param.Tags)
}

func (svc *Service) CountNews(param *CountNewsReq) (int64, error) {
	return svc.dao.CountNews(param.Title)
}

func (svc *Service) ListNews(param *ListNewsReq, pager *app.Pager) ([]*dao.News, error) {
	return svc.dao.ListNews(param.Title, pager.PageIndex, pager.PageSize)
}
