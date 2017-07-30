package controllers

import (
	"gisbo/models"

	"github.com/astaxie/beego"
)

//
type LoginController struct {
	beego.Controller
}

// login method
func (this *LoginController) Get() {
	this.Layout = "layout/layout.html"
	this.TplName = "login/login.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "login/css.tpl"
	this.LayoutSections["Scripts"] = "login/scripts.tpl"
	this.LayoutSections["Sidebar"] = ""
}

func (this *LoginController) Signin() {
	flash := beego.NewFlash()
	username, password := this.Input().Get("username"), this.Input().Get("password")
	user, err := models.SignIn(username, password)
	//response := NewResponse()
	if err == nil {
		this.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)

		//this.Response = make(map[string]interface{})
		//this.Response
		response := SuccessResponse("登录成功！", &user)
		//response.Success = true
		//response.Message = "登录成功！"
		//response.Data["data"] = &user
		this.Data["json"] = &response
		//this.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&this.Controller)
		/*
			response.Success = false
			response.Message = "用户名不存在或密码错误！"
			response.Data["err"] = &err
		*/
		response := FailResponse("用户名不存在或密码错误！", &err)
		this.Data["json"] = &response
		//c.Redirect("/login", 302)
	}

	this.ServeJSON()
}
