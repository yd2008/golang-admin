package errcode

var (
	CountNewsFail   = NewError(30010001, "获取新闻总数失败")
	GetNewsListFail = NewError(30010002, "获取新闻列表失败")
	CreateNewsFail  = NewError(30010003, "创建新闻失败")
)
