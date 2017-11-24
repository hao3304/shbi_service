package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Log struct {
	Id int `orm:"pk;auto"`
	Mail string
	Action string
	Content string
	Created int64
}

func (this * Log)GetAll(page int) (interface{},error) {
	o := orm.NewOrm()
	var logs []*Log
	o.Raw(fmt.Sprintf("select * from log ORDER BY created DESC  limit %d , 20",page)).QueryRows(&logs)
	count,err := o.QueryTable(new(Log)).Count()
	result := make(map[string]interface{})

	result["Data"] = logs
	result["Total"] = count
	result["Page"] = page

	return result,err
}

func (this *Log)Add() (int64, error) {
	o := orm.NewOrm()
	fmt.Println(this)
	this.Created = time.Now().Unix()
	return o.Insert(this)
}

func init()  {
	orm.RegisterModel(new(Log))
}

