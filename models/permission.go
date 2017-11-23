package models

import "github.com/astaxie/beego/orm"

type Permission struct {
	Id int `orm:"pk;auto"`
	Name string
	/**
	0 目录 1 菜单 2 路径
	 */
	//Type int
	//Url string `orm:"null"`
	Pk string
	Roles []*Role `orm:"reverse(many)"`
	//Icon string
}

func init()  {
	orm.RegisterModel(new(Permission))
}