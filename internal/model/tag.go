package model

import (
	"github.com/go-sql-driver/mysql"
	"golang-admin/pkg/errcode"
	"gorm.io/gorm"
	"strings"
)

type Tag struct {
	Common
	Title    string `json:"title" gorm:"column:title;unique;type:varchar(64);comment:标签标题"`
	IsEnable uint8  `json:"is_enable" gorm:"column:is_enable;default:1;type:int(2);comment:是否可用"`
}

func (Tag) TableName() string {
	return "go_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(Tag{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) Create(db *gorm.DB) error {
	err := db.Create(&t).Error
	if sqlErr, ok := err.(*mysql.MySQLError); ok {
		err = errcode.CreateTagFail.WithDetails(sqlErr.Error())
	}
	return err
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}

func (t Tag) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&t).Updates(values).Error
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	err := db.First(&tag, t.ID).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	err := db.Offset(pageOffset).Limit(pageSize).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func queryTags(db *gorm.DB, tagStr string) []*Tag {
	var tags []*Tag
	tagsArr := strings.Split(tagStr, ",")
	err := db.Table("go_tag").Find(&tags, tagsArr).Error
	if err != nil {
		// TODO: 2021/04/30 错误需要写入日志
		return nil
	}

	return tags
}