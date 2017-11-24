package controllers


type UserController struct {
	BaseController
}

func (this *UserController)Get()  {
	username := this.Ctx.Input.GetData("username")
	this.Success(username)
}

