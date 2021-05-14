package v1

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/dao"
	"golang-admin/internal/model"
	"golang-admin/internal/service"
	"golang-admin/pkg/app"
	"golang-admin/pkg/convert"
	"golang-admin/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

/*
@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	路由，从左到右分别为：路由地址，HTTP 方法
*/

// Register godoc
// @Summary 注册用户
// @Tags 用户
// @Produce json
// @Param data body service.LoginUserBody true "注册信息"
// @Success 200 {object} model.SwagSuccess
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register [post]
func (User) Register(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.RegisterUserBody{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.UserRegister(&param); err != nil {
		response.ErrorIfHasDetail(err, errcode.RegisterUserFail)
		return
	}

	response.Success()
}

// Login godoc
// @Summary 用户登录
// @Tags 用户
// @Produce json
// @Param data body service.LoginUserBody true "登录信息"
// @Success 200 {object} swagUserLogin "登录成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func (User) Login(c *gin.Context) {
	response := app.NewResponse(c)
	var param = service.LoginUserBody{}
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

	token, err := app.GenerateToken(user.ID)
	if err != nil {
		response.Error(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.SuccessData(token)
}

// Update godoc
// @Summary 用户信息更新
// @Tags 用户
// @Produce json
// @Param id path int true "用户id"
// @Param data body service.UpdateUserBody true "更新信息"
// @Success 200 {object} model.SwagSuccess "更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/users/{id} [put]
func (User) Update(c *gin.Context) {
	response := app.NewResponse(c)
	var id = convert.Str2uint(c.Param("id"))
	if id == 0 {
		response.Error(errcode.IdRequiredError)
		return
	}

	var param = service.UpdateUserBody{}
	if validErrors, ok := app.BindAndValid(c, &param); !ok {
		response.Error(errcode.InvalidParams.WithDetails(validErrors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UserUpdate(id, &param)
	if err != nil {
		response.ErrorIfHasDetail(err, errcode.UpdateUserFail)
		return
	}

	response.Success()
}

// Get godoc
// @Summary 获取用户
// @Tags 用户
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} swagUser "获取用户成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/users/{id} [get]
func (User) Get(c *gin.Context) {
	response := app.NewResponse(c)
	var id = convert.Str2uint(c.Param("id"))
	if id == 0 {
		response.Error(errcode.IdRequiredError)
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.UserGet(id)
	if err != nil {
		response.ErrorIfHasDetail(err, errcode.GetUserFail)
		return
	}

	response.SuccessData(user)
}

// Delete godoc
// @Summary 删除用户
// @Tags 用户
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} model.SwagSuccess "删除用户成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/users/{id} [delete]
func (User) Delete(c *gin.Context) {
	param := service.DeleteUserReq{ID: convert.Str2uint(c.Param("id"))}
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

type swagUser struct {
	model.SwagCommon
	Data *dao.User `json:"data"`
}

type swagUserLogin struct {
	model.SwagCommon
	Data *swagToken `json:"data"`
}

type swagToken struct {
	Token string `json:"token"`
}
