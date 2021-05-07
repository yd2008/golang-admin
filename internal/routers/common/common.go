package common

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/model"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

type Common struct {}

func NewCommon() Common {
	return Common{}
}

// GetOssAccessToken godoc
// @Summary 获取oss凭证
// @Tags 通用
// @Produce json
// @Success 200 {object} swagOssCredentials "获取oss凭证成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /common/getossaccesstoken [get]
func (Common) GetOssAccessToken(c *gin.Context) {
	response := app.NewResponse(c)
	credentials, err := app.CreateCredentials()
	if err != nil {
		response.Error(errcode.OssInternalError)
		return
	}

	response.SuccessData(credentials)
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

type swagOssCredentials struct {
	model.SwagCommon
	Data *app.OssCredentials `json:"data"`
}
