package main

import (
	_ "shbi_service/routers"
	_ "shbi_service/models"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default","sqlite3","./db/sqlite.db")
	orm.RunCommand()
}

func main() {
	beego.Run()
}

