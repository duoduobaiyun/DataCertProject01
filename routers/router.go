package routers

import (
	"DataCertProject01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//用户注册的接口请求
	beego.Router("/user_register",&controllers.RegisterController{})

	//直接登录的页面请求接口
	beego.Router("/login.html",&controllers.LoginController{})

	//用户登录请求接口
	beego.Router("/user_login",&controllers.LoginController{})

	//用户上传文件
	beego.Router("/", &controllers.UploadController{},"*:Get")
	beego.Router("/home", &controllers.UploadController{},"*:Post")
}
