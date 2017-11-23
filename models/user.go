package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int `orm:"pk;auto"`
	UserName string `orm:"unique"`
	Password string
	Roles []*Role `orm:"rel(m2m)"`
	Created int64
}



func init()  {
	orm.RegisterModel(new(User))
}

