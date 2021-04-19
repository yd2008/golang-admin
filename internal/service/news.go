package service

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/app"
)

type CountNewsReq struct {
	Title string `json:"title" binging:"max=100"`
}

type ListNewsReq struct {
	Title string `form:"title"`
}

func (svc *Service) CountNews(param *CountNewsReq) (int64, error) {
	return svc.dao.CountNews(param.Title)
}

func (svc *Service) ListNews(param *ListNewsReq, pager *app.Pager) ([]*model.News, error) {
	return svc.dao.ListNews(param.Title, pager.PageIndex, pager.PageSize)
}
