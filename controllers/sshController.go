package controllers

import (
	"shbi_service/utrl"
	"fmt"
)

type SshController struct {
	BaseController
}

func (this *SshController)Get() {
	mail := this.GetString("mail")
	action := this.GetString("action")
	rep,err := doAction(action,mail)
	if err != nil {
		this.Success(rep)
	}else{
		this.Fail("执行错误",500)
	}
}

func doAction(action string,mail string) (string, error)  {
	ssh := new(utrl.SshCMD)
	ssh.LoadPEM("/tmp/id_rsa")
	var cmd string
	switch action {
	case "locked":
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ma %s zimbraAccountStatus locked;/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail,mail)
	case "active":
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ma %s zimbraAccountStatus active;/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail,mail)
	default:
		cmd = fmt.Sprintf("/opt/zimbra/bin/zmprov ga %s zimbraAccountStatus",mail)
	}

	return ssh.RemoteRun("nicstaff","202.121.179.34",cmd)


}