package controllers

import (
	"fmt"
	_ "strconv"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

type User struct {
	Username string   `form:"username" json:"uname"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
}

func (c *ArticleController) Get() {

	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("id参数错误!")
		return
	}
	fmt.Printf("%T---%v", id, id)
	c.Ctx.WriteString("新闻列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("添加新闻-----")
}

func (c *ArticleController) EditArticle() {

	username := c.GetString("username")
	password := c.GetString("password", "你的密码呢")
	id, _ := c.GetInt("id")
	c.Ctx.WriteString("编辑新闻-----" + username + "---" + password)

	u := User{}
	c.ParseForm(&u)
	fmt.Printf("%T@@@@@%#v", u, u)
	c.Data["json"] = u
	c.ServeJSON()
	fmt.Println(id)

}
