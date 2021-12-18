package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gin-swagger-demo/pkg/gredis"
	"gin-swagger-demo/pkg/logging"
	"gin-swagger-demo/pkg/setting"
	"gin-swagger-demo/pkg/util"
	"gin-swagger-demo/routers"
	"gin-swagger-demo/models"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
	models.Setup()
}
// @title gin-swagger-demo
// @version 1.0
// @description gin-swagger-demo服务后端API接口文档
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name wanggang
// @license.url https://gitee.com/wanggang826/gin-swagger-demo
// @host 127.0.0.1:8000
// @BasePath
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
