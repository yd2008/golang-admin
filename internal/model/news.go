package model

import (
	"gorm.io/gorm"
)

type News struct {
	Common
	Title      string `json:"title" gorm:"column:title;type:varchar(100);comment:标题"`
	Content    string `json:"content" gorm:"column:content;type:varchar(255);comment:内容"`
	Tags       string `json:"tags" gorm:"column:tags;type:varchar(100);comment:新闻标签"`
}

type NewsInfo struct {
	News *News
	Tags []*Tag
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

func (n News) Create(db *gorm.DB) error {
	return db.Create(&n).Error
}

func (n News) List(db *gorm.DB, offset, limit int) ([]*NewsInfo, error) {
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
	var newsDaos []*NewsInfo
	for _, _news := range news {
		var newsDao = NewsInfo{
			News: _news,
			Tags: queryTags(db, _news.Tags),
		}

		newsDaos = append(newsDaos, &newsDao)
	}
	return newsDaos, nil
}
