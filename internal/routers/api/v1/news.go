package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/dao"
	"golang-admin/internal/model"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type News struct{}

func NewNews() News {
	return News{}
}

// Create godoc
// @Summary 新建新闻
// @Tags 新闻
// @Produce json
// @Param data body service.CreateNewsBody true "创建新闻"
// @Success 200 {object} model.SwagSuccess
// @Failure 400 {object} errcode.Error
// @Failure 500 {object} errcode.Error
// @Router /api/v1/news [post]
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

// List godoc
// @Summary 新闻列表
// @Tags 新闻
// @Produce json
// @Param param query swagListQuery true "新闻列表"
// @Success 200 {object} swagListRes
// @Failure 400 {object} errcode.Error
// @Failure 500 {object} errcode.Error
// @Router /api/v1/news [get]
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

type swagListQuery struct {
	model.SwagPager
	Title string `json:"title"`
}

type swagListRes struct {
	model.SwagCommon
	Pager *app.Pager `json:"pager"`
	List  []*dao.News
}
