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
	return Response{ctx:ctx}
}

func (r *Response) Success() {
	r.SuccessData(nil)
}

func (r *Response) SuccessData(data interface{}) {
	data = gin.H{
		"code": 0,
		"data": data,
		"msg" : "请求成功",
	}
	r.ctx.JSON(http.StatusOK, data)
}

func (r *Response) Error(err *errcode.Error) {
	h := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		h["details"] = details
	}
	r.ctx.JSON(err.StatusCode(), h)
}
