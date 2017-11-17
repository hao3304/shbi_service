package models

import "github.com/astaxie/beego/orm"

type Log struct {
	Id int `orm:"pk;auto"`
	Mail string
	Action string
	Created int64
}

func GetAll() []Log {
	o:= orm.NewOrm()
	logs := []Log{}
	o.QueryTable("log").Limit(100).All(&logs)
	return logs
}

func init()  {
	orm.RegisterModel(new(Log))
}

