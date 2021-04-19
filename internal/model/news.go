package model

import (
	"gorm.io/gorm"
)

type News struct {
	Common
	Title      string `json:"title" gorm:"column:title;type:varchar(100);comment:标题"`
	Content    string `json:"content" gorm:"column:content;type:varchar(255);comment:内容"`
}

func (News) TableName() string {
	return "go_news"
}

func (n News) Count(db *gorm.DB) (int64, error) {
	var count int64
	var err error
	if n.Title != "" {
		db = db.Where("title LIKE ?", "%"+n.Title+"%")
	}
	if err = db.Model(&News{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (n News) List(db *gorm.DB, offset, limit int) ([]*News, error) {
	var news []*News
	var err error
	if n.Title != "" {
		db = db.Where("title Like ?", "%"+n.Title+"%")
	}
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if err = db.Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}
