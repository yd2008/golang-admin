package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type News struct{}

func NewNews() News {
	return News{}
}

func (n News) Create(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.CreateNewsBody{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	var svc = service.New(c.Request.Context())
 	err := svc.NewsCreate(&param)
	if err != nil {
		response.Error(errcode.CreateNewsFail)
		return
	}

	response.Success()
}


func (n News) List(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.ListNewsReq{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	pager := app.Pager{PageIndex: app.GetPageIndex(c), PageSize: app.GetPageSize(c)}
	var svc = service.New(c.Request.Context())
	totalSize, err := svc.CountNews(&service.CountNewsReq{Title: param.Title})
	if err != nil {
		response.Error(errcode.CountNewsFail)
		return
	}

	list, err := svc.ListNews(&param, &pager)
	if err != nil {
		response.Error(errcode.GetNewsListFail)
		return
	}

	response.SuccessList(list, pager, totalSize)
}
