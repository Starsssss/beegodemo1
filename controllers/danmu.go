package controllers

import (
	"fmt"
	 "strconv"
	"beegodemo1/models"

	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


// func init(){
// 	//注册定义的model
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
//     orm.RegisterDataBase("default", "mysql", "root:20080808@/movie?charset=utf8")
// }

type DanmuController struct {
	beego.Controller
}



// type Danmu struct {
// 	id   int
// 	create_time string
// 	update_time string
// 	movie_id int
// 	content string
// 	second int
// }


/**
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "movie___id": 1,
            "content": "1的弹幕1",
            "second": 1
        },
        {
            "id": 2,
            "movie___id": 1,
            "content": "1的2222",
            "second": 2
        },
        {
            "id": 3,
            "movie___id": 1,
            "content": "电影1的弹幕3",
            "second": 3
        },
        {
            "id": 4,
            "movie___id": 1,
            "content": "电影1的弹幕4",
            "second": 4
        }
    ]
}


 ["1.0000000",0,16777215,"1","弹幕内容2"],
*/

 type result_list struct{
	Content string
	Second int
	Abc string 
}

func (c *DanmuController) SaveDanmu() {
	// movie_id, err := c.GetInt("movie_id")
	c.Data["json"] = c.GetString("content")//弹幕内容
	c.Data["json"] = c.GetString("second")//弹幕在第几秒
	second,_:=strconv.Atoi(c.GetString("second"))
	movie_id,_:= c.GetInt("movie_id",0)
	content:=c.GetString("content")
fmt.Println("********",second,"***",content,"_______",c.GetString("second"))
	
	err:=models.SaveDanmu(second, content,movie_id)
	resMap :=make(map[string]interface{})
	if(err==nil){
		
	// countryCapitalMap :=make(map[string]string)
	resMap [ "code" ] = 0
	resMap [ "msg" ]  = "发送弹幕成功!"
	fmt.Println(resMap)
	c.Data["json"] = resMap
	
	}else{
		resMap [ "code" ] = 1
		resMap [ "msg" ]  = "发送弹幕失败!"
	}


	c.ServeJSON()
}


func (c *DanmuController) GetDanmu() {
	movie_id, err := c.GetInt("movie_id")
	if err!=nil{
		fmt.Println(err)
	}
	


	res:=models.GetDemoOneByMovieId(movie_id)
	slice1 := make([] models.Danmu,len(res))
	
	// c.Data["json"] = res
	slice1=res
	fmt.Println("----------",slice1)




	
	
	// list_slice:= make([] interface{},5)
	arr:=[5] interface{} {1,2,3,4,"5"}
	var rowss []interface{}
	for i := 0; i < len(res); i++ {
		arr [0]=strconv.Itoa(res[i].Second)+".0000000"
		arr [1]=res[i].Second
		arr [2]=16777215
		arr [3]="1"
		arr [4]=res[i].Content
		rowss=append(rowss,arr)
	 }
	 fmt.Println("@@@@@@@@@@@@",rowss)
	
	countryCapitalMap :=make(map[string]interface{})
	// countryCapitalMap :=make(map[string]string)
	countryCapitalMap [ "code" ] = 0
	countryCapitalMap [ "data" ] = rowss
	fmt.Println(countryCapitalMap)
	c.Data["json"] = countryCapitalMap
	c.ServeJSON()
	fmt.Println(res)
}

