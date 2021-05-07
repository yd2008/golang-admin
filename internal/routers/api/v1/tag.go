package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/model"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// Create godoc
// @Summary 新建标签
// @Tags 标签
// @Produce json
// @Param data body service.CreateTagBody true "创建标签"
// @Success 200 {object} model.SwagSuccess
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (Tag) Create(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.CreateTagBody{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.TagCreate(&param)
	if err != nil {
		response.ErrorIfHasDetail(err, errcode.CreateTagFail)
		return
	}

	response.Success()
}

// List godoc
// @Summary 标签列表
// @Tags 标签
// @Produce json
// @Success 200 {object} swagTagList
// @Failure 400 {object} errcode.Error
// @Failure 500 {object} errcode.Error
// @Router /api/v1/tags [get]
func (Tag) List(c *gin.Context) {
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	totalCount, err := svc.TagCount()
	if err != nil {
		response.Error(errcode.CountTagFail)
		return
	}
	tags, err := svc.TagList()
	if err != nil {
		response.ErrorIfHasDetail(err, errcode.CreateTagFail)
		return
	}

	response.SuccessListAll(tags, totalCount)
}

type swagTagList struct {
	model.SwagTotalSize
	List []*model.Tag `json:"list"`
}
