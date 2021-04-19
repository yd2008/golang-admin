package dao

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/app"
)

func (d *Dao) CountNews(title string) (int64, error) {
	var news = model.News{Title: title}
	return news.Count(d.engine)
}

func (d *Dao) ListNews(title string, pageIndex, pageSize int) ([]*model.News, error) {
	var news = model.News{Title: title}
	offset := app.GetPageOffset(pageIndex, pageSize)
	return news.List(d.engine, offset, pageSize)
}
