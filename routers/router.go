package routers

import (
	"shbi_service/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("service/ssh", &controllers.SshController{})
    beego.Router("service/mode/?:id", &controllers.ModeController{})
}
