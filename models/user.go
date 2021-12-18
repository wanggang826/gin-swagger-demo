package models

import "gin-swagger-demo/pkg/logging"

type User struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Openid          string `xorm:"not null default '' comment('微信openid') VARCHAR(32)"`
	OpenidXchx      string `xorm:"not null default '' comment('小程序openid') index VARCHAR(32)"`
	Unionid         string `xorm:"not null default '' comment('微信unionId') VARCHAR(32)"`
	SessionKey      string `xorm:"not null default '' comment('session_key') VARCHAR(32)"`
	Subscribe       int    `xorm:"not null default 0 comment('是否关注：0=否，1=是') TINYINT(1)"`
	Nickname        string `xorm:"not null default '' comment('微信昵称') VARCHAR(50)"`
	Sex             int    `xorm:"not null default 0 comment('用户的性别，1=男性，2=女性，0=未知') TINYINT(1)"`
	City            string `xorm:"not null default '' comment('用户所在城市') VARCHAR(20)"`
	Country         string `xorm:"not null default '' comment('用户所在国家') VARCHAR(20)"`
	Province        string `xorm:"not null default '' comment('用户所在省份') VARCHAR(20)"`
	Language        string `xorm:"not null default '' comment('用户的语言，简体中文为zh_CN') VARCHAR(20)"`
	Headimgurl      string `xorm:"not null default '' comment('微信头像') VARCHAR(256)"`
	PhoneNumber     string `xorm:"not null default '' comment('手机号，带区号') VARCHAR(16)"`
	PurePhoneNumber string `xorm:"not null default '' comment('手机号，没区号') VARCHAR(16)"`
	CountryCode     string `xorm:"not null default '' comment('国家码') VARCHAR(8)"`
	SubscribeTime   int    `xorm:"not null default 0 comment('用户关注时间,如果用户曾多次关注，则取最后关注时间') INT(10)"`
	SubscribeScene  string `xorm:"not null default '' comment('返回用户关注的渠道来源') VARCHAR(64)"`
	Email           string `xorm:"not null default '' comment('邮箱') VARCHAR(50)"`
	IsBanned        int `xorm:"not null default 0 comment('是否被禁用') TINYINT(1)"`
	CreateTime      int    `xorm:"created not null default 0 comment('创建时int间') INT(10)"`
	UpdateTime      int    `xorm:"updated not null default 0 comment('最后更新时间') INT(10)"`
}

func (t *User) Add() (int64, error) {
	res, err := db.Insert(t)
	if err != nil {
		logging.Error("DB: User.Add: ", "msg",err.Error())
	}
	return res, err
}

func (item *User) Get() (bool, error) {
	res, err := db.Get(item)
	if err != nil {
		logging.Error("DB: User.Get: ", "msg",err.Error())
	}
	return res, err
}

func (item *User) Edit(closEdit string) (int64, error) {
	res, err := db.ID(item.Id).Cols(closEdit).Update(item)
	if err != nil {
		logging.Error("DB: User.Edit: ", "msg",err.Error())
	}
	return res, err
}

func GetUserByOpenid(openid string) *User {
	user := &User{Openid: openid}
	has, err := db.Get(user)
	if err != nil {
		logging.Error("DB: User.GetUserByOpenid: ", "msg",err.Error())
		return nil
	}

	if !has {
		return nil
	}
	return user
}

func GetUserById(id int) *User {
	user := &User{Id: id}
	has, err := db.Get(user)
	if err != nil {
		logging.Error("DB: User.GetUserById: ", "msg",err.Error())
		return nil
	}
	if !has {
		return nil
	}
	return user
}

func GetUserByNickName(nickName string) *User {
	user := &User{Nickname: nickName}
	has, err := db.Get(user)
	if err != nil {
		logging.Error("DB: User.GetUserByNickName: ", "msg",err.Error())
		return nil
	}

	if !has {
		return nil
	}
	return user
}

func GetAllUserSearchCount(keyword string,status string) (int, error) {
	whereSql := ""
	if keyword != "" {
		whereSql += "nickname like '%" + keyword + "%'"
	}
	if status != "" {
		whereSql += "AND is_banned = " + status
	}

	item := new(User)
	total, err := db.Table("user").Where(whereSql).Count(item)
	if err != nil {
		logging.Error("DB: team.GetAllUserSearchCount Count: ", "msg",err.Error())
	}

	return int(total), err
}

func GetsUserByNickname(keyword string) ([]*User, error) {
	whereSql := ""
	if keyword != "" {
		whereSql += " nickname like '%" + keyword + "%'"
	}

	ls := make([]*User, 0)
	err := db.Table("user").Where(whereSql).OrderBy("id desc").Find(&ls)
	if err != nil {
		logging.Error("DB: Team.GetsUserByNickname: ", "msg",err.Error())
	}
	return ls, err
}

func GetAllUserSearchPage(keyword string,status string,limit int, offset int) ([]*User, error) {
	whereSql := ""
	if keyword != "" {
		whereSql += " nickname like '%" + keyword + "%'"
	}

	if status != "" {
		whereSql += "AND is_banned = " + status
	}

	ls := make([]*User, 0)
	err := db.Table("user").Where(whereSql).Limit(limit, offset).OrderBy("id desc").Find(&ls)
	if err != nil {
		logging.Error("DB: Team.GetAllUserSearchPage: ", "msg",err.Error())
	}

	return ls, err
}

func GetTodayUserCount(todayTime int) (int,error) {
	item := new(User)
	count, err := db.Table("user").Where("create_time >= ?",todayTime).Count(item)
	if err != nil {
		logging.Error("DB: User.GetTodayUserCount Count: ", "msg",err.Error())
	}
	return int(count), err
}

func GetTotalUserCount() (int,error) {
	item := new(User)
	count, err := db.Table("user").Count(item)
	if err != nil {
		logging.Error("DB: User.GetTotalUserCount Count: ", "msg",err.Error())
	}
	return int(count), err
}

func GetUserCountByTime(startTime int,endTime int) (int,error) {
	item := new(User)
	count, err := db.Table("user").Where("create_time >= ? and create_time <= ?",startTime,endTime).Count(item)
	if err != nil {
		logging.Error("DB: User.GetUserCountByTime Count: ", "msg",err.Error())
	}
	return int(count), err
}