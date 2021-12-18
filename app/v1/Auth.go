package v1

import (
	"gin-swagger-demo/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"

	"gin-swagger-demo/pkg/app"
	"gin-swagger-demo/pkg/e"
	"gin-swagger-demo/pkg/util"
)

func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := appG.GetUid()

	token, err := util.GenerateUserToken(userId)
	if err != nil {
		logging.Error("RefreshToken err:","msg",err.Error(),"userId",userId)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}