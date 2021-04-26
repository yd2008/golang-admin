package app

import (
	"github.com/gin-gonic/gin"
	"golang-admin/global"
	"golang-admin/pkg/convert"
)

type Pager struct {
	PageIndex int   `json:"page_index"`
	PageSize  int   `json:"page_size"`
	TotalPage int64 `json:"total_page"`
	TotalSize int64 `json:"total_size"`
}

func GetPageIndex(c *gin.Context) int {
	pageIndex := convert.Str2int(c.Query("page_index"))
	if pageIndex <= 0 {
		return 1
	}

	return pageIndex
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.Str2int(c.Query("page_size"))
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(pageIndex, pageSize int) int {
	var offset = 0
	if pageIndex > 0 {
		offset = (pageIndex - 1) * pageSize
	}

	return offset
}
