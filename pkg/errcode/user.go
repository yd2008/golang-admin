package errcode

var (
	RegisterUserFail    = NewError(20010001, "注册用户失败")
	LoginUserFail       = NewError(20010002, "账号或者密码错误")
	DeleteUserFail      = NewError(20010003, "删除用户失败")
	GetUserFail         = NewError(20010004, "获取用户失败")
)
