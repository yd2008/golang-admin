package service

import "golang-admin/internal/model"

type CreateTagBody struct {
	Title    string `json:"title" binding:"required,min=2,max=10"`
	IsEnable uint8  `json:"is_enable" binding:"oneof=0 1"`
}

func (svc *Service) TagCount() (int64, error) {
	return svc.dao.CountTag()
}

func (svc *Service) TagCreate(param *CreateTagBody) error {
	return svc.dao.CreateTag(param.Title, param.IsEnable)
}

func (svc *Service) TagList() ([]*model.Tag, error) {
	return svc.dao.ListTag()
}
