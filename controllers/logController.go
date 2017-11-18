package controllers

import (
	"shbi_service/models"
)

type LogController struct {
	BaseController
}

func (this *LogController)Get() {
	page, err := this.GetInt("page")
	if err !=nil {
		this.Fail("参数错误",400)
	}
	log := new(models.Log)
	result,err := log.GetAll(page)
	if err != nil {
		this.Fail("参数错误",400)
	}
	this.Success(result)
}


