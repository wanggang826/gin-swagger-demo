package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-swagger-demo/pkg/e"
	"gin-swagger-demo/pkg/logging"
	"gin-swagger-demo/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = e.SUCCESS
		var data interface{}

		uid,isAdmin,err := util.GetUidFromHeader(c)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			logging.Error("jwt check token error: ","msg", err.Error(),"uid",uid,"isAdmin",isAdmin)
		} else {
			if !isAdmin && c.Request.URL.Path[:7] == "/admin/"{
				logging.Error("jwt ILLEGAL request : ", uid)
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				c.Set("uid", uid)
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
