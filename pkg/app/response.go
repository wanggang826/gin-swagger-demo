package app

import (
	"github.com/gin-gonic/gin"

	"gin-swagger-demo/models"
	"gin-swagger-demo/pkg/e"
)

type Gin struct {
	C    *gin.Context
	uid  int
	user *models.User
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) GetUser() *models.User {
	if nil == g.user {
		uid := g.GetUid()
		g.user = models.GetUserById(uid)
	}
	return g.user
}

// GetUid 获取当前登录用户的 UID
func (g *Gin) GetUid() int {
	if g.uid <= 0 {
		g.uid = g.C.GetInt("uid")
	}
	return g.uid
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
