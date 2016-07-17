package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	"app/models"
	"github.com/astaxie/beego/orm"
)

//初始化数据库
func init()  {
	models.RegisterDB()
}

func main() {
	//默认情况下,不会建表,所以需要自动建表
	orm.Debug = true	//orm开启debug
	//自动建表	默认数据库  是否每次都重新建表:false  打印数据库信息:true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

