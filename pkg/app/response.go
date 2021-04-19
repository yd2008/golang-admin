package app

import (
	"github.com/gin-gonic/gin"
	"golang-admin/pkg/errcode"
	"net/http"
)

type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) Response {
	return Response{ctx: ctx}
}

func (r *Response) Success() {
	r.SuccessData(nil)
}

func (r *Response) SuccessData(data interface{}) {
	data = gin.H{
		"code": 0,
		"data": data,
		"msg":  "请求成功",
	}
	r.ctx.JSON(http.StatusOK, data)
}

func (r *Response) SuccessList(list interface{}, pager Pager, totalSize int64) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			PageIndex: GetPageIndex(r.ctx),
			PageSize:  GetPageSize(r.ctx),
			TotalSize: totalSize,
		},
	})
}

func (r *Response) Error(err *errcode.Error) {
	h := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		h["details"] = details
	}
	r.ctx.JSON(err.StatusCode(), h)
}

// ErrorIfHasDetail 如果已经是自定义错误则返回自定义错误，
//                  如不是可以在第二个参数设置需要返回的自定义错误
func (r *Response) ErrorIfHasDetail(err error, errCode *errcode.Error) {
	detailErr, ok := err.(*errcode.Error)
	if ok {
		r.Error(detailErr)
		return
	}

	if err != nil {
		r.Error(errCode)
		return
	}
}
