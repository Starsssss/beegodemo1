package models

import (  
    "time"

    "github.com/astaxie/beego"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/garyburd/redigo/redis"
)
var (
    Db *gorm.DB
    Redis redis.Conn
    err error
)

func init(){
    mysqlInit()
    Redis = PoolConnect()
}

//初始化mysql数据库
func mysqlInit(){
     dbhost := beego.AppConfig.String("dbhost")
     dbport := beego.AppConfig.String("dbport")
     dbuser := beego.AppConfig.String("dbuser")
     dbpassword := beego.AppConfig.String("dbpassword")
     dbname := beego.AppConfig.String("dbname")
     //设置数据库默认端口
     if dbport == ""{
        dbport = "3306"
     }
     // root:root@tcp(127.0.0.1:3306)/beego_blog?charset=utf8&parseTime=True&loc=Local
     url := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname +"?charset=utf8&parseTime=True&loc=Local"
     Db, err = gorm.Open("mysql", url)
     // 单数表
     Db.SingularTable(true)
     // 自动建表
    //  Db.AutoMigrate(
    //         &Comment{},



    //     )
     // 在main.go 中延迟关闭连接
}

//直接连接
func Connect() redis.Conn {
    pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
    return pool
}

//通过连接池
func PoolConnect() redis.Conn {
    // 建立连接池
    pool := &redis.Pool{
        MaxIdle:     5000,              //最大空闲连接数
        MaxActive:   10000,             //最大连接数
        IdleTimeout: 180 * time.Second, //空闲连接超时时间
        Wait:        true,              //超过最大连接数时，是等待还是报错
        Dial: func() (redis.Conn, error) { //建立链接
            c, err := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
            if err != nil {
                return nil, err
            }
            // 选择db
            //c.Do("SELECT", '')
            return c, nil
        },
    }
    return pool.Get()
}