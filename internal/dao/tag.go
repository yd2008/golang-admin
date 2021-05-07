package dao

import (
	"golang-admin/internal/model"
)

func (d *Dao) CountTag() (int64, error) {
	return model.Tag{}.Count(d.engine)
}

func (d *Dao) CreateTag(title string, isEnable uint8) error {
	var tag = model.Tag{
		Title:    title,
		IsEnable: isEnable,
	}
	return tag.Create(d.engine)
}

func (d *Dao) DeleteTag(id uint) error {
	var tag = model.Tag{
		Common: model.Common{
			ID: id,
		},
	}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTag(id uint) (*model.Tag, error) {
	var tag = model.Tag{
		Common: model.Common{
			ID: id,
		},
	}
	return tag.Get(d.engine)
}

func (d *Dao) UpdateTag(id uint, title string, isEnable uint8) error {
	var tag = model.Tag{
		Common: model.Common{
			ID: id,
		},
	}
	var updateInfo = make(map[string]interface{})
	updateInfo["title"] = title
	updateInfo["is_enable"] = isEnable
	return tag.Update(d.engine, updateInfo)
}

func (d *Dao) ListTag() ([]*model.Tag, error) {
	return model.Tag{}.List(d.engine, 0, 999)
}

func getTagsDao(tags []*model.Tag) []string {
	var tagsDao []string
	if len(tags) == 0 {
		return tagsDao
	}
	for _, tag := range tags {
		tagsDao = append(tagsDao, tag.Title)
	}
	return tagsDao
}