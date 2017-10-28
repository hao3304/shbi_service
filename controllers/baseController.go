package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type Rep struct {
	Code int
	Message string
	Response interface{}
}

func (this *BaseController)Success(data interface{}) {
	rep := Rep{
		Code:0,
		Message:"OK",
		Response:data,
	}
	this.Data["json"] = &rep
	this.ServeJSON()
}

func (this *BaseController)Fail(msg string,code int)  {
	rep := Rep{
		Code: code,
		Message: msg,
	}
	this.Data["json"] = &rep
	this.ServeJSON()
}

func (this *BaseController)CheckToken(token string)  {

}