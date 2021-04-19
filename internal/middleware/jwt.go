package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang-admin/pkg/app"
	"golang-admin/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errCode = errcode.Success
		tokenStr, exist := c.GetQuery("token")
		if !exist {
			tokenStr = c.GetHeader("token")
		}
		if tokenStr == "" {
			errCode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(tokenStr)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					errCode = errcode.UnauthorizedTokenTimeout
				default:
					errCode = errcode.UnauthorizedTokenError
				}
			}
		}

		if errCode != errcode.Success {
			response := app.NewResponse(c)
			response.Error(errCode)
			c.Abort()
			return
		}

		c.Next()
	}
}
