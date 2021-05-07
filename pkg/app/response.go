package app

import (
	"github.com/gin-gonic/gin"
	"golang-admin/pkg/errcode"
	"math"
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
		"msg":  "操作成功！",
	}
	r.ctx.JSON(http.StatusOK, data)
}

func (r *Response) SuccessList(list interface{}, pager Pager, totalSize int64) {
	pager.TotalSize = totalSize
	pager.TotalPage = int64(math.Ceil(float64(totalSize) / float64(pager.PageSize)))
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"list": list,
		"pager": pager,
		"msg":  "操作成功！",
	})
}

func (r *Response) SuccessListAll(list interface{}, totalSize int64) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"list": list,
		"total_size": totalSize,
		"msg":  "操作成功！",
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

// ErrorIfHasDetail
// 如果已经是自定义错误则返回自定义错误，
// 如不是可以在第二个参数设置需要返回的自定义错误
func (r *Response) ErrorIfHasDetail(err error, errCode *errcode.Error) {
	detailErr, ok := err.(*errcode.Error)
	if ok {
		r.Error(detailErr)
		return
	}

	r.Error(errCode)
}
