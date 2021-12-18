package routers

import (
	v1 "gin-swagger-demo/admin/v1"
	"gin-swagger-demo/middleware/jwt"
)

/*InitAdminRouter
初始化 admin 接口路由
*/
func InitAdminRouter() {
	adminv1 := r.Group("/admin/v1")
	adminv1.POST("/auth/login", v1.Login)

	adminv1.Use(jwt.JWT())
	{
		//refresh token
		adminv1.GET("/auth/refreshToken", v1.RefreshToken)
		adminv1.POST("/auth/changePassword", v1.ChangePassword)

		//help
		adminv1.GET("/help/getHelpList", v1.GetHelpPage)
		adminv1.POST("/help/addHelp", v1.AddHelp)
		adminv1.POST("/help/editHelp", v1.EditHelp)
		adminv1.POST("/help/delHelp", v1.DelHelp)
	}
}
