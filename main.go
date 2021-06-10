package main

import (
	_"fmt"
	_ "beegodemo1/routers"
	_"beegodemo1/models"
	_"net/http"
_	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	logs.SetLogger("console")
	// logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)

	logs.EnableFuncCallDepth(true)




	orm.RegisterDriver("mysql", orm.DRMySQL)
	

	orm.RegisterDataBase("default", "mysql", "root:20080808@/movie?charset=utf8")
}



func main() { 


	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        //允许访问所有源
		AllowAllOrigins: true,
        //可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
        //其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"}, 
        //指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
        //公开的HTTP标头列表
		ExposeHeaders:	[]string{"Content-Length"},
        //如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))



	//  err:= http.ListenAndServe("127.0.0.1:8000", nil)

	// if err != nil {
    //     fmt.Println("ListenAndServe: ", err) 
    // }


	beego.Run()
	
}
