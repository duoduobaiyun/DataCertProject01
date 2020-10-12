package controllers

import (
	sha2562 "crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get()  {
	this.TplName="home.html"
}

func (this *UploadController)Post()  {
	file,head,err:=this.GetFile("file")
	if err != nil {
		this.Ctx.WriteString("获取文件失败")
		return
	}
	defer file.Close()

	filename:=head.Filename
	sha256:=sha2562.New()
	sha256.Write([]byte(filename))
	byte:=sha256.Sum(nil)
	filename=hex.EncodeToString(byte)
	err=this.SaveToFile("file","static/img/"+filename)
	if err != nil {
		this.Ctx.WriteString("上传失败")
	}else {
		this.Ctx.WriteString("上传成功")
	}
	this.TplName="home.html"
}