package main

import (
	_ "DataCertProject01/routers"
	"github.com/astaxie/beego"
	"DataCertProject01/db_mysql"
)

func main() {
	//连接数据库
    db_mysql.ConnectDB()
	//beego.SetStaticPath("/js","./static/js")
	//beego.SetStaticPath("/css","./static/css")
	//beego.SetStaticPath("/img","./static/img")
	beego.Run()

}

