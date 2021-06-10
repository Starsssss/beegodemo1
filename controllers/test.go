package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}


func (c *TestController) Hello() {

	c.Ctx.WriteString("欢迎来到Beego!")
}