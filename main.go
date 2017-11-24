package main

import (
	_ "shbi_service/routers"
	_ "shbi_service/models"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"shbi_service/models"
)

func init()  {
	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default","sqlite3","./db/sqlite.db")
	orm.RunCommand()
	//ssh := new(util.SshCMD)
	//ssh.LoadPEM("/tmp/id_rsa")
	//rep,err:=ssh.RemoteRun("nicstaff","202.121.179.34","/opt/zimbra/bin/zmprov ga aquarius@sjtu.edu.cn zimbraAccountStatus locked")

	//if err !=nil {
	//	 fmt.Println(err)
	//}
	//fmt.Println(rep)

	fmt.Println("start.....")
	o := orm.NewOrm()
	user := models.User{UserName:"admin"}
	err := o.Read(&user, "UserName")

	if err == orm.ErrNoRows {
		fmt.Println("end")
		user.Password = "hao123456"
		user.Created = time.Now().Unix()
		o.Insert(&user)
	}else{
		fmt.Println("123")
	}
}

func main() {
	beego.Run()
}

