package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.html"
	c.Data["isHome"] = true
	c.Data["title"] = "beego 博客 - 首页"
	c.Data["isLogin"] = checkAccount(c.Ctx)

}
