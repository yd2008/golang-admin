package model

import (
	"gorm.io/gorm"
	"time"
)

type Common struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Empty struct{}

type SwagSuccess struct {
	SwagCommon
}

type SwagTotalSize struct {
	SwagCommon
	TotalSize int `json:"total_size"`
}

type SwagCommon struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type SwagPager struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
