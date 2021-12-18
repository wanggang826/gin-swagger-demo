package models

import "gin-swagger-demo/pkg/logging"

type Admin struct {
	Id            int    `xorm:"not null pk autoincr comment('主键ID') INT(11)"`          //主键ID
	Username      string `xorm:"not null default '' comment('用户名') CHAR(20)"`           //用户名
	Password      string `xorm:"not null default '' comment('登录密码') VARCHAR(80)"`       //登录密码
	Salt          string `xorm:"default '' comment('随机码') CHAR(16)"`                    //随机码
	IsActivited   int    `xorm:"not null default 0 comment('是否激活：0=否，1=是') TINYINT(4)"` //是否激活：0=否，1=是
	ActivatedTime int    `xorm:"not null default 0 comment('激活时间') INT(10)"`            //激活时间
	AdminType     int    `xorm:"not null comment('1管理员 2超级管理员') TINYINT(4)"`            //1管理员 2超级管理员
	Permissions   string `xorm:"comment('权限，json存储') TEXT"`                             //权限，json存储
	CreateTime    int    `xorm:"created not null default 0 comment('创建时间') INT(10)"`    //创建时间
	UpdateTime    int    `xorm:"updated not null default 0 comment('最后更新时间') INT(10)"`  //最后更新时间
}

func (item *Admin) Add() (int64, error) {
	res, err := db.Insert(item)
	if err != nil {
		logging.Error("DB: Admin.Add: ", "msg", err.Error())
	}
	return res, err
}

func (item *Admin) Edit(closEdit string) (int64, error) {
	res, err := db.ID(item.Id).Cols(closEdit).Update(item)
	if err != nil {
		logging.Error("DB: Admin.Edit: ", "msg", err.Error())
	}
	return res, err
}

func (item *Admin) Delete() (int64, error) {
	res, err := db.ID(item.Id).Delete(item)
	if err != nil {
		logging.Error("DB: Admin.Delete: ", "msg", err.Error())
	}
	return res, err
}

func (item *Admin) GetById() (bool, error) {
	res, err := db.ID(item.Id).Get(item)
	if err != nil {
		logging.Error("DB: Admin.GetById: ", "msg", err.Error())
	}
	return res, err
}

func (item *Admin) Get() (bool, error) {
	res, err := db.Get(item)
	if err != nil {
		logging.Error("DB: Admin.Get: ", "msg", err.Error())
	}
	return res, err
}
