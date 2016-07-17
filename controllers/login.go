package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController)Get() {
	c.Data["isCategory"] = true
	c.Data["title"] = "beego 博客 - 登陆"
	c.TplName = "login.html"
}

func (c *LoginController)Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autoLogin := c.Input().Get("auto")

	if ( username != beego.AppConfig.String("username") ||
	password != beego.AppConfig.String("password")) {
		c.Redirect("/login",301)
		return
	}

	maxAge := 0
	if autoLogin == "1"{
		maxAge = 1<<31 -1
	}

	c.Ctx.SetCookie("username",username,maxAge,"/");
	c.Ctx.SetCookie("password",password,maxAge,"/");

	c.Redirect("/",301)
	return
}

func checkAccount(cxt *context.Context) bool {

	//beego.Info(cxt.GetCookie("username"))
	if cxt.GetCookie("username") != beego.AppConfig.String("username") &&
	 cxt.GetCookie("password") != beego.AppConfig.String("password"){
		 return false
	 }
	//fmt.Println(cxt.GetCookie("username"))
	//fmt.Println(cxt.GetCookie("password"))
	//log := logs.Logger()
	//log.SetLogger("console")
	//l := logs.GetLogger()
	//beego.Debug(cxt.GetCookie("username"))
	return true
}


