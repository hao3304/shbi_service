package main

import (
	_ "shbi_service/routers"
	_ "shbi_service/models"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
	"shbi_service/utrl"
	"fmt"
)

func init()  {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default","sqlite3","./db/sqlite.db")
	orm.RunCommand()

	ssh := new(utrl.SshCMD)
	ssh.LoadPEM("/tmp/id_rsa")
	rep,err:=ssh.RemoteRun("nicstaff","202.121.178.34","pwd")

	if err !=nil {
		 fmt.Println(err)
	}
	fmt.Println(rep)
}

func main() {
	beego.Run()
}

