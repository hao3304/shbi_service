package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id int `orm:"pk;auto"`
	Name string
	Users []*User `orm:"reverse(many)"`
	Permissions []*Permission `orm:"rel(m2m)"`
}

func init()  {
	orm.RegisterModel(new(Role))
}
