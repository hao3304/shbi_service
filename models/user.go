package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type User struct {
	Id int `orm:"pk;auto"`
	Name string
	UserName string `orm:"unique"`
	Password string
	Visible bool
	Roles []*Role `orm:"rel(m2m)"`
	Created int64
}


func (this * User)GetAll(page int) (interface{},error) {
	o := orm.NewOrm()
	var users []*User
	o.Raw(fmt.Sprintf("select * from user where user_name != 'admin' ORDER BY created DESC  limit %d , 20",page)).QueryRows(&users)
	count,err := o.QueryTable(new(User)).Count()
	result := make(map[string]interface{})

	result["Data"] = users
	result["Total"] = count
	result["Page"] = page

	return result,err
}

func (this *User)Add() (int64, error)  {

	o := orm.NewOrm()
	this.Created = time.Now().Unix()
	return o.Insert(this)

}

func init()  {
	orm.RegisterModel(new(User))
}

