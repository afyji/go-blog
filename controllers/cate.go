package controllers

import (
	"github.com/astaxie/beego"
	"app/models"
)

type CateController struct {
	beego.Controller
}

func (c *CateController)Get() {
	c.TplName = "cate.html"
	c.Data["isCate"] = true
	c.Data["title"] = "beego 博客 - 分类"
	c.Data["isLogin"] = checkAccount(c.Ctx)

	op := c.Input().Get("op")

	switch op {
	case "add":
		name := c.Input().Get("name")
		if name == ""{
			break
		}
		err := models.AddCate(name)
		if err != nil {
			beego.Error(err)
		}
	case "del":
		id := c.Input().Get("id")
		beego.Info(id)
		if id == ""{
			break
		}
		err := models.DelCate(id)
		if err != nil{
			beego.Error(err)
		}
	}
	cates,err := models.GetAllCate()
	if err != nil{
		beego.Error(err)
		c.Redirect("/category",301)
		return
	}

	beego.Info(cates)
	c.Data["cates"] = cates
}

