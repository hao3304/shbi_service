package routers

import (
	"shbi_service/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("service/mode/?:id", &controllers.ModeController{})
}
