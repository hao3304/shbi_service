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
	fmt.Println("start ssh")
	ssh := new(utrl.SshCMD)
	rsa := ssh.LoadPEM("/tmp/id_rsa")
	fmt.Println(rsa)
	rep,err:=ssh.RemoteRun("nicstaff","202.121.179.34","pwd")

	if err !=nil {
		 fmt.Println(err)
	}
	fmt.Println(rep)
}

func main() {
	beego.Run()
}

