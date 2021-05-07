package errcode

var (
	CreateTagFail = NewError(30000001, "新建标签失败")
	CountTagFail  = NewError(30000002, "获取标签总数失败")
)
