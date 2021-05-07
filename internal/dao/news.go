package dao

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/app"
)

type News struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (d *Dao) CountNews(title string) (int64, error) {
	var news = model.News{Title: title}
	return news.Count(d.engine)
}

func (d *Dao) CreateNews(title, content, tags string) error {
	var news = model.News{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	return news.Create(d.engine)
}

func (d *Dao) ListNews(title string, pageIndex, pageSize int) ([]*News, error) {
	var news = model.News{Title: title}
	offset := app.GetPageOffset(pageIndex, pageSize)
	list, err := news.List(d.engine, offset, pageSize)
	if err != nil {
		return nil, err
	}
	var newsArr []*News
	for _, info := range list {
		newsArr = append(newsArr, News{}.createNewsDao(info))
	}
	return newsArr, nil
}

func (n News) createNewsDao(info *model.NewsInfo) *News {
	var _newsDao = News{
		ID:      info.News.Common.ID,
		Title:   info.News.Title,
		Content: info.News.Content,
		Tags:    getTagsDao(info.Tags),
	}
	return &_newsDao
}

