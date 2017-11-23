package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"shbi_service/models"
	"shbi_service/util"
	"fmt"
)

type LoginController struct {
	BaseController
}


func (this *LoginController) Post() {
	o := orm.NewOrm()
	inputs := new(models.User)

	json.Unmarshal(this.Ctx.Input.RequestBody,&inputs)
	user := models.User{
		UserName:inputs.UserName,
	}
	fmt.Println(inputs)

	err := o.Read(&user,"UserName")
	if err==nil {
		if user.Password == inputs.Password {
			t, expires := util.GenToken(&user)
			token := struct{
				Token string
				Expires int64
			}{
				Token:t,
				Expires:expires,
			}
			this.Success(token)
		}else{
			this.Fail("用户名或者密码错误!",10001)
		}

	}else {
		this.Fail("用户名或者密码错误!",10001)
	}
}