package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

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
