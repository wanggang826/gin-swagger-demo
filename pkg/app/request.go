package app

import (
	"github.com/astaxie/beego/validation"

	"gin-swagger-demo/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error("MarkErrors","key",err.Key, "msg",err.Message)
	}

	return
}
