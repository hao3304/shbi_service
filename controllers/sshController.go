package controllers

import (
	"shbi_service/util"
	"fmt"
	"github.com/astaxie/beego/logs"
	"shbi_service/models"
	"time"
)

type SshController struct {
	BaseController
}

func (this *SshController)Get() {

	mail := this.GetString("mail")
	action := this.GetString("action")
	rep,err := doAction(action,mail)
	if err == nil {
		auth := this.Ctx.Input.GetData("username")
		log := new(models.Log)
		log.Created = time.Now().Unix()
		log.Mail = mail
		log.Action = action
		log.Content = rep
		log.UserName = auth.(string)
		log.Add()

		this.Success(rep)
	}else{
		logs.Error(err)
		this.Fail("执行错误",500)
	}
}

func doAction(action string,mail string) (string, error)  {
	ssh := new(util.SshCMD)
	ssh.LoadPEM("/tmp/id_rsa")
	var cmd string
	switch action {
	case "locked":
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ma %s zimbraAccountStatus locked;/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail,mail)
	case "active":
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ma %s zimbraAccountStatus active;/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail,mail)
	case "gmi":
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov gmi %s ",mail)
	default:
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail)
	}
	logs.Info("cmd: %s",cmd)
	return ssh.RemoteRun("nicstaff","202.121.179.34",cmd)


}