package controllers

import (
	"shbi_service/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController)Get()  {

	auth := this.Ctx.Input.GetData("username")
	if auth != "admin" {
		this.Ctx.Abort(403,"没有权限")
	}

	page, err := this.GetInt("page")
	if err !=nil {
		this.Fail("参数错误",400)
	}
	user := new(models.User)
	result,err := user.GetAll(page)
	if err != nil {
		this.Fail("参数错误",400)
	}
	this.Success(result)
}

func (this *UserController)Post() {

	auth := this.Ctx.Input.GetData("username")
	if auth != "admin" {
		this.Fail("没有权限",403)
	}

	user := new(models.User)
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	_,err := user.Add()
	if err != nil {
		beego.Debug(err)
		this.Fail(err.Error(),400)
	}

	this.Success(&user)
}


func (this *UserController)Delete() {

	auth := this.Ctx.Input.GetData("username")
	if auth != "admin" {
		this.Ctx.Abort(403,"没有权限")
	}

	var id int
	this.Ctx.Input.Bind(&id, ":id")
	if id != 0 {
		user := new(models.User)
		user.Id = id
		if num, err := orm.NewOrm().Delete(user);err == nil&&num>0 {
			this.Success("删除成功！")
		}else{
			this.Fail("没有相关记录",400)
		}
	}else{
		this.Fail("参数错误！",400)
	}
}

func (this *UserController)Patch()  {
	var id int
	this.Ctx.Input.Bind(&id, ":id")

	if id != 0 {
		userId  := this.Ctx.Input.GetData("userId")
		strId := strconv.Itoa(id)
		if userId == strId {
			o := orm.NewOrm()
			user := new(models.User)
			user.Id = id
			json.Unmarshal(this.Ctx.Input.RequestBody, &user)
			if user.Password == "" {
				this.Ctx.Abort(400,"Password不能为空")
			}

			if num,err :=o.Update(user,"Password"); num>0&&err ==nil {
				this.Success("更新成功")
			}else{
				this.Fail(err.Error(),400)
			}

		}else{
			this.Ctx.Abort(403,"没有权限")
		}

	}else{
		auth := this.Ctx.Input.GetData("username")
		if auth != "admin" {
			this.Ctx.Abort(403,"没有权限")
		}

		o := orm.NewOrm()
		user := new(models.User)
		json.Unmarshal(this.Ctx.Input.RequestBody, &user)
		if user.Id == 0 {
			this.Ctx.Abort(400,"Id不能为空")
		}
		if user.Name == "" {
			this.Ctx.Abort(400,"Name不能为空")
		}
		if user.UserName == "" {
			this.Ctx.Abort(400,"UserName不能为空")
		}
		if user.Password == "" {
			this.Ctx.Abort(400,"Password不能为空")
		}

		if num,err :=o.Update(user,"Name","Password","Visible","UserName"); num>0&&err ==nil {
			this.Success("更新成功")
		}else{
			this.Fail(err.Error(),400)
		}
	}
}