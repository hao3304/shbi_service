package controllers

import (
	"github.com/astaxie/beego/orm"
	"shbi_service/models"
	"encoding/json"
	"fmt"
	"time"
)

type ModeController struct {
	BaseController
}

func (this *ModeController)Get()  {
	var modes []*models.Mode
	var id int
	this.Ctx.Input.Bind(&id, ":id")

	if id !=0 {
		mode := models.Mode{Id:id}
		orm.NewOrm().Read(&mode)
		this.Success(&mode)
	}else{
		orm.NewOrm().QueryTable("mode").All(&modes)
		this.Success(&modes)
	}


}

func (this *ModeController)Post()  {
	o := orm.NewOrm()
	mode := models.Mode{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &mode)

	if mode.Name == "" {
		this.Fail("Name不能为空",400)
	}
	if mode.Query == "" {
		this.Fail("Query不能为空",400)
	}
	t := time.Now()
	mode.Created = t.Unix()

	_,err := o.Insert(&mode)
	if err == nil {
		this.Success(&mode)
	}else{
		fmt.Println(err)
		this.Fail("参数错误",400)
	}
}

func (this *ModeController)Delete() {
	var id int
	this.Ctx.Input.Bind(&id, ":id")
	if id != 0 {
		if num, err := orm.NewOrm().Delete(&models.Mode{Id:id});err == nil&&num>0 {
			this.Success("删除成功！")
		}else{
			this.Fail("没有相关记录",400)
		}
	}else{
		this.Fail("参数错误！",400)
	}
}

func (this *ModeController)Patch()  {
	o := orm.NewOrm()
	mode := models.Mode{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &mode)

	if mode.Id == 0 {
		this.Fail("Id不能为空",400)
	}
	if mode.Name == "" {
		this.Fail("Name不能为空",400)
	}
	if mode.Query == "" {
		this.Fail("Query不能为空",400)
	}

	if num,err :=o.Update(&mode,"Name","Query","Index","Visible"); num>0&&err ==nil {
		this.Success("更新成功")
	}else{
		fmt.Println(err)
		this.Fail("更新错误，没有相关记录。",400)
	}


}