package models

import (
	"gin-swagger-demo/pkg/logging"
)

type Help struct {
	Id         int    `xorm:"not null default '' comment('主键ID') pk autoincr"`
	Title      string `xorm:"not null default '' comment('标题') VARCHAR(256)"`
	H5Link     string `xorm:"not null default '' comment('H5链接地址') VARCHAR(256) 'H5_LINK'"`
	Sort       int    `xorm:"not null default 0 comment('排序 大的在前') INT(10) "`
	CreateTime int    `xorm:"created not null default 0 comment('创建时间') INT(10) "`
}

func (item *Help) Add() (int64, error) {
	res, err := db.Insert(item)
	if err != nil {
		logging.Error("DB: Help.Add: ", "msg",err.Error())
	}
	return res, err
}

func (item *Help) Edit(closEdit string) (int64, error) {
	res, err := db.ID(item.Id).Cols(closEdit).Update(item)
	if err != nil {
		logging.Error("DB: Help.Edit: ", "msg",err.Error())
	}
	return res, err
}

func (item *Help) Delete() (int64, error) {
	res, err := db.Delete(item)
	if err != nil {
		logging.Error("DB: Help.Delete: ","msg", err.Error())
	}
	return res, err
}

func (item *Help) GetById() (bool, error) {
	res, err := db.ID(item.Id).Get(item)
	if err != nil {
		logging.Error("DB: Help.GetById: ", "msg",err.Error())
	}
	return res, err
}

func (item *Help) Get() (bool, error) {
	res, err := db.Get(item)
	if err != nil {
		logging.Error("DB: Help.Get: ", "msg",err.Error())
	}
	return res, err
}

func GetHelpList(limit int, offset int) ([]*Help, error) {
	ls := make([]*Help, 0)
	err := db.Table("help").Where("").Limit(limit, offset).OrderBy("sort desc").Find(&ls)
	if err != nil {
		logging.Error("DB: GetHelpList: ", "msg",err.Error())
	}
	return ls, err
}

func GetHelpCount() (int, error) {
	item := new(Help)
	count, err := db.Table("help").Where("").Count(item)
	return int(count), err
}