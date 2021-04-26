package common

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type Common struct {}

func NewCommon() Common {
	return Common{}
}

func (Common) WechatLogin(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.WehatLoginBody{}

	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.WechatLogin(&param)
	if err != nil {
		response.Error(errcode.WechatLoginError)
		return
	}
	response.SuccessData(user)
}
