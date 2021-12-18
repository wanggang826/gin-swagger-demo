package v1

import (
	"gin-swagger-demo/models"
	"gin-swagger-demo/pkg/app"
	"gin-swagger-demo/pkg/e"
	"gin-swagger-demo/pkg/logging"
	"gin-swagger-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary 列表查询
// @Description 帮助文档分页列表查询
// @Tags 帮助文档
// @Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param pageSize query int false "每页条数"
// @Param pageNo query int false "页码"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} app.Response {"code":400,"data":null,"msg":""}
// @Router /admin/v1/help/getHelpList [get]
func GetHelpPage(c *gin.Context) {
	appG := app.Gin{C: c}

	//分页参数
	limit := com.StrTo(c.Query("pageSize")).MustInt()
	pageNo := com.StrTo(c.Query("pageNo")).MustInt()
	offset := 0
	if pageNo > 0 {
		offset = (pageNo - 1) * limit
	}

	total,_ := models.GetHelpCount()
	helpList, _ := models.GetHelpList(limit, offset)
	type responseParams struct {
		Id         int
		Title      string
		H5Link     string
		Sort       int
		CreateTime string
	}
	retList := make([]responseParams, len(helpList))
	for k, v := range helpList {
		retList[k] = responseParams{
			Id:         v.Id,
			Sort:       v.Sort,
			Title:      v.Title,
			H5Link:     v.H5Link,
			CreateTime: util.TimeToStrDate(v.CreateTime),
		}
	}
	result := util.ResponsePageData(total, retList, limit, pageNo)
	appG.Response(http.StatusOK, e.SUCCESS, result)
}

type HelpSubmitParams struct {
	Id int
	Title string //标题
	H5Link string //链接
	Sort int //排序
}

// @Summary 新增
// @Description 新增帮助文档
// @Tags 帮助文档
// @Security Bearer
// @Produce  json
// @Param help body HelpSubmitParams true "Title：标题；H5Link 链接；sort 排序"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}
// @Router /admin/v1/help/addHelp [post]
func AddHelp (c *gin.Context){
	appG := app.Gin{C: c}
	var params HelpSubmitParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Error("GetRequest: Gin BindJSON","msg",err.Error(),"params",params)
		appG.Response(http.StatusOK, e.ERROR_COMMON_PARAM_GIN_BINDJSON_ERROR, nil)
		return
	}

	valid := validation.Validation{}
	if v := valid.Required(params.Title, "Title"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}
	if v := valid.Required(params.H5Link, "H5Link"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}

	help := &models.Help{
		Title: params.Title,
		H5Link: params.H5Link,
		Sort: params.Sort,
	}
	help.Add()
	appG.Response(http.StatusOK, e.SUCCESS, help.Id)
}

// @Summary 修改
// @Description 修改帮助文档
// @Tags 帮助文档
// @Security Bearer
// @Produce  json
// @Param help body HelpSubmitParams true "Title：标题；H5Link 链接；sort 排序"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}
// @Router /admin/v1/help/editHelp [post]
func EditHelp(c *gin.Context) {
	appG := app.Gin{C: c}
	var params HelpSubmitParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Error("GetRequest: Gin BindJSON","msg",err.Error(),"params",params)
		appG.Response(http.StatusOK, e.ERROR_COMMON_PARAM_GIN_BINDJSON_ERROR, nil)
		return
	}

	valid := validation.Validation{}
	if v := valid.Required(params.Id, "Id"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}
	if v := valid.Required(params.Title, "Title"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}
	if v := valid.Required(params.H5Link, "H5Link"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}
	help := &models.Help{
		Id: params.Id,
		Title: params.Title,
		H5Link: params.H5Link,
		Sort: params.Sort,
	}
	_,edit_err := help.Edit("title,h5_link,sort")
	if edit_err != nil {
		logging.Error("help edit error","msg",edit_err.Error(),"helpId",help.Id,"params",params)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, help.Id)
}

type DelParams struct {
	Id int //帮助文档ID
}

// @Summary 删除
// @Description 帮助文档删除
// @Tags 帮助文档
// @Security Bearer
// @Produce  json
// @Param id path int true "ID"
// @Param help body DelParams true "id:帮助文档ID"
// @Success 200 {object} app.Response {"code":200,"data":null,"msg":""}
// @Router /admin/v1/help/delHelp [post]
func DelHelp(c *gin.Context)  {
	appG := app.Gin{C: c}
	var params HelpSubmitParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Error("GetRequest: Gin BindJSON","msg",err.Error(),"params",params)
		appG.Response(http.StatusOK, e.ERROR_COMMON_PARAM_GIN_BINDJSON_ERROR, nil)
		return
	}

	valid := validation.Validation{}
	if v := valid.Required(params.Id, "Id"); !v.Ok {
		appG.Response(http.StatusOK, e.ERROR_COMMON_URL_PARAM_ERROR, nil)
		return
	}
	help := &models.Help{Id: params.Id}
	help.Delete()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}