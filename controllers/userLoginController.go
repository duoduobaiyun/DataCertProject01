package controllers

import (
	"DataCertProject01/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
   l.TplName="login.html"
}

/*
用户登录接口
*/
func (l *LoginController) Post() {
	var user models.User
	err:=l.ParseForm(&user)
	if err != nil {
		//l.TplName="error.html"
		l.Ctx.WriteString("抱歉，用户数据解析失败，请重试！")
		return
	}
	//查询数据库的用户信息
	u,err := user.QueryUser()
	if err != nil {
		//l.TplName="error.html"
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登入失败，请重试！")
		return
	}
	//登入成功,跳转项目核心功能页面(home.html)
	l.Data["Phone"] = u.Phone
	l.TplName="home.html"//{{.Phone}}

}