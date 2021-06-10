package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["Website"] = "你好啊@@@"
	c.Data["Email"] = "astaxie@gmail.com%%%%%%%%%%"
	c.TplName = "goods.tpl"
}
