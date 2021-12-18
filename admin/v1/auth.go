package v1

import (
	"gin-swagger-demo/models"
	"gin-swagger-demo/pkg/logging"
	"github.com/astaxie/beego/validation"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-swagger-demo/pkg/app"
	"gin-swagger-demo/pkg/e"
	"gin-swagger-demo/pkg/util"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// @Summary 登录
// @Description 管理员登录
// @Tags 管理员
// @Security Bearer
// @Produce  json
// @Param admin body LoginParams true "Username:用户名,Password:密码"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}
// @Router /admin/v1/auth/login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	var params LoginParams
	err := c.BindJSON(&params)
	logging.Info("auth Login params", params)
	if err != nil {
		logging.Error("GetRequest: Login Gin BindJSON:", "msg", err.Error())
		appG.Response(http.StatusOK, e.ERROR_COMMON_PARAM_GIN_BINDJSON_ERROR, nil)
		return
	}

	valid := validation.Validation{}
	valid.Required(params.Username, "Username")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.ERROR_ADMIN_LOGIN_USERNAME_EMPTY, nil)
		return
	}

	valid.Required(params.Password, "Password")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	admin := &models.Admin{
		Username: params.Username,
	}

	exist, _ := admin.Get()
	if !exist {
		logging.Info("auth change exist", "===============1======", params.Password)
		appG.Response(http.StatusOK, e.ERROR_ADMIN_LOGIN_USERNAME_PASSWORD_ERROR, nil)
		return
	}

	dbSalt := admin.Salt
	checkPass := util.Md5String(params.Password + dbSalt)
	if checkPass != admin.Password {
		logging.Info("auth  Password diff", "===============2======", params.Password)
		appG.Response(http.StatusOK, e.ERROR_ADMIN_LOGIN_USERNAME_PASSWORD_ERROR, nil)
		return
	}

	salt := util.RandomString(16)

	newPassword := util.Md5String(params.Password + salt)
	admin.Salt = salt
	admin.Password = newPassword
	closEdit := "salt,password"
	admin.Edit(closEdit)

	token, err := util.GenerateAdminToken(admin.Id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}

func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}
	adminId := appG.GetUid()

	token, err := util.GenerateAdminToken(adminId)
	if err != nil {
		logging.Error("error auth token:", "msg", err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}

type ChangePasswordParams struct {
	Password   string
	RePassword string
}

// @Summary 修改密码
// @Description 管理员修改密码
// @Tags 管理员
// @Security Bearer
// @Produce  json
// @Param admin body ChangePasswordParams true "Password 密码 RePassword 再次输入密码"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}
// @Router /admin/v1/auth/changePassword [post]
func ChangePassword(c *gin.Context) {
	appG := app.Gin{C: c}
	adminId := appG.GetUid()
	var params ChangePasswordParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Error("GetRequest: Login Gin BindJSON:", "msg", err.Error())
		appG.Response(http.StatusOK, e.ERROR_COMMON_PARAM_GIN_BINDJSON_ERROR, nil)
		return
	}
	valid := validation.Validation{}
	valid.Required(params.Password, "Password")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	valid.Required(params.RePassword, "RePassword")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if params.Password != params.RePassword {
		appG.Response(http.StatusOK, e.ERROR_ADMIN_CHANGE_PASSWORD_DIFF_ERROR, nil)
		return
	}

	admin := &models.Admin{
		Id: adminId,
	}
	exist, _ := admin.Get()
	if !exist {
		appG.Response(http.StatusOK, e.ERROR_USER_GETINFO_FAIL, nil)
		return
	}
	logging.Info("auth change password", "password", params.Password)
	salt := util.RandomString(16)

	newPassword := util.Md5String(params.Password + salt)
	logging.Info("auth change password", "newPassword", newPassword)
	admin.Salt = salt
	admin.Password = newPassword
	closEdit := "salt,password"
	admin.Edit(closEdit)

	token, err := util.GenerateAdminToken(admin.Id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, token)
}
