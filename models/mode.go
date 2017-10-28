package models

import (
	"github.com/astaxie/beego/orm"
)

type Mode struct {
	Id int `orm:"pk;auto"`
	Name string
	Query string
	Visible bool
	Index int
	Created int64
	//Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init()  {
	orm.RegisterModel(new(Mode))
}