package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	InitRegisterModel()
}

/*
初始化时注册模型
*/
func InitRegisterModel() {
	orm.RegisterModel(new(Danmu))
	fmt.Printf("%v", GetDemoOneByMovieId)
	// GetDemoOneById(1)
}

/*
--------------------------------
使用CURD操作数据库，需要注册模型
orm.RegisterModel(new(Demo))
--------------------------------
*/

type Danmu struct {
	Id          int    `orm:"column(id);pk" json:"id,omitempty"` // 设置主键
	create_time string `json:"create___time,omitempty"`
	update_time string `json:"update___time,omitempty"`
	Movie_id    int    `json:"movie___id,omitempty"`
	Content     string `json:"content,omitempty"`
	Second      int    `json:"second,omitempty"`
}

// func (Danmu *Danmu) String() string{
// 	res:="@@"+string(Danmu.Id)+" ==="+Danmu.Content
// 	return res
// }

/*
通过主键id查询
*/
func GetDemoOneByMovieId1(id int) Danmu {
	Danmu := Danmu{Movie_id: id} //Id: id,
	err := orm.NewOrm().Read(&Danmu, "Movie_id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("##########", Danmu.Id, Danmu.Content)
	}
	return Danmu
}

func GetDemoOneByMovieId(id int) []Danmu {
	//https://www.cnblogs.com/hei-ma/articles/13716916.html
	o := orm.NewOrm()
	// QueryTable使用方式一：按表名查询（将表名当参数传给QueryTable方法）
	// qs := o.QueryTable("danmu")
	Danmu_arr :=[] Danmu{} //Id: id,

	
	

	//All：所有数据         返回符合查询条件的所有数据，最多可返回1000行数据，如果超过一千行则只返回一千行。
	num,err:=o.QueryTable("danmu").Filter("movie_id__exact", id).All(&Danmu_arr,"content","second")
	fmt.Println("---------",num,err)
	return Danmu_arr
}
func SaveDanmu(second int,content string,movie_id int) error {

	o := orm.NewOrm()
	var danmu Danmu
	danmu.Content = content
	danmu.Second = second
	danmu.Movie_id = movie_id
	
	id, err := o.Insert(&danmu)
	if err == nil {
		fmt.Println(id)
	}
	return err
}
