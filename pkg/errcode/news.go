package errcode

var (
	CountNewsFail   = NewError(30010001, "获取新闻总数失败")
	GetNewsListFail = NewError(30010002, "获取新闻列表失败")
	//DeleteUserFail   = NewError(20010003, "删除用户失败")
)
