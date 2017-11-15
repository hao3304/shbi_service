package models

import "github.com/astaxie/beego/orm"

type Log struct {
	Id int `orm:"pk;auto"`
	Mail string
	Action string
	Created int64
}


func init()  {
	orm.RegisterModel(new(Log))
}

