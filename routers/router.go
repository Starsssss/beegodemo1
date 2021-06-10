package routers

import (
	"beegodemo1/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/goods", &controllers.GoodsController{})
	// beego.Router("/article", &controllers.ArticleController{})
	// beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	// beego.Router("/article/edit", &controllers.ArticleController{}, "post:EditArticle")



	
	beego.Router("/danmu", &controllers.DanmuController{}, "get:GetDanmu")
	beego.Router("/danmu/GetDanmu", &controllers.DanmuController{}, "get:GetDanmu")
	beego.Router("/danmu/send", &controllers.DanmuController{}, "post:SaveDanmu")//接收发送的弹幕
	// beego.Router("/", &controllers.DanmuController{}, "get:GetDanmu")
	beego.Router("/", &controllers.TestController{}, "get:Hello")

}
