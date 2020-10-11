package routers

import (
	"DataCertProject01/controllers"
	"DataCertProject01/db_mysql"
	"github.com/astaxie/beego"
)

func init() {
	db_mysql.ConnectDB()
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})
    beego.Router("/login.html",&controllers.LoginController{})
	beego.Router("/user_login",&controllers.LoginController{})
}
