package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"gin-swagger-demo/pkg/logging"
	"gin-swagger-demo/pkg/setting"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Setup() {
	ConnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	engine, err := xorm.NewEngine("mysql", ConnStr)
	db = engine

	if err != nil {
		logging.Error("Fail to create xorm system logger: ", err.Error())
	} else {

		err := db.Ping()
		if err != nil {
			logging.Error("Connect to mysql error: ", err.Error())
		} else {
			logging.Info("Connect to mysql OK : ", ConnStr)
		}
	}

	if setting.ServerSetting.RunMode == "debug" {
		db.ShowSQL(true)
	}

}

