package util

import (
	"gin-swagger-demo/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"math"
)

// GetPage get page parameters
func GetAppPageParams(c *gin.Context) (limit int, offset int) {
	limit = com.StrTo(c.Query("limit")).MustInt()
	offset = com.StrTo(c.Query("offset")).MustInt()
	if limit < 1 {
		limit = constants.DEFAULT_PAGE_LIMIT
	}

	return
}

// GetPage get page parameters
func GetPageParams(c *gin.Context) (limit int, offset int, pageNo int) {
	offset = 0
	limit = com.StrTo(c.Query("pageSize")).MustInt()
	pageNo = com.StrTo(c.Query("pageNo")).MustInt()
	if pageNo > 0 {
		offset = (pageNo - 1) * limit
	}
	if limit < 1 {
		limit = constants.DEFAULT_PAGE_LIMIT
	}

	return
}

func ResponsePageData(total int, data interface{}, pageSize int, pageNo int) map[string]interface{} {
	var responseMap = make(map[string]interface{})

	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))

	responseMap["data"] = data
	responseMap["pageSize"] = pageSize
	responseMap["pageNo"] = pageNo
	responseMap["totalPage"] = totalPage
	responseMap["totalCount"] = total

	return responseMap
}
