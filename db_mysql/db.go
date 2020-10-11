package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
var Db *sql.DB
func ConnectDB(){
	//、读取config配置信息
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//、组织连接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp +")/" + dbName + "?charset=utf8"
	//、连接数据库
	db,err := sql.Open(dbDriver,connUrl)
	//、获取数据库连接对象，处理连接结果
	if err !=nil {
		panic("数据库连接错误，请检查配置。")
	}
	Db = db //为全局赋值
}
