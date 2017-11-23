package routers

import (
	"shbi_service/controllers"
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/context"
	"shbi_service/util"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    ns := beego.NewNamespace("service",
			beego.NSBefore(func(ctx *context.Context) {
				authToken := ctx.Input.Header("Authorization")
				beego.Debug("auth token:", authToken)
				kv := strings.Split(authToken, " ")
				if len(kv) != 2 || kv[0] != "Bearer" {
					beego.Error("AuthString invalid:", authToken)
					ctx.Abort(503,"认证信息格式错误。")
				}
				token := kv[1]
				claim, err :=util.ParseJwt(token)
				if err == nil {
					ctx.Input.SetData("username",claim.Audience)
				}
				ctx.Abort(503,"认证失败，请重新登陆。")
			}),
    		beego.NSRouter("/ssh", &controllers.SshController{}),
    		beego.NSRouter("/log", &controllers.LogController{}),
    		beego.NSRouter("/login", &controllers.LoginController{}),
    		beego.NSRouter("/mode/?:id", &controllers.ModeController{}),
    	)
    	beego.AddNamespace(ns)



}
