package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/convert"
	"golang-admin/pkg/errcode"
)

type User struct {}

func NewUser() User {
	return User{}
}

func (User) Register(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.RegisterUserReq{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UserRegister(&param)
	if err != nil {
		response.Error(errcode.RegisterUserFail)
		return
	}

	response.Success()
}

func (User) Login(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.LoginUserReq{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.UserLogin(&param)
	if err != nil {
		response.Error(errcode.LoginUserFail)
		return
	}

	response.SuccessData(user)
}

func (User) Get(c *gin.Context) {
	param := service.GetUserReq{convert.Str2uint(c.Param("id"))}
	response := app.NewResponse(c)
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.UserGet(&param)
	if err != nil {
		response.Error(errcode.DeleteUserFail)
		return
	}

	response.SuccessData(user)
}

func (User) Delete(c *gin.Context) {
	param := service.DeleteUserReq{ID:convert.Str2uint(c.Param("id"))}
	response := app.NewResponse(c)
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UserDelete(&param)
	if err != nil {
		response.Error(errcode.DeleteUserFail)
		return
	}

	response.Success()
}
