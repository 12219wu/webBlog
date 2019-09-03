package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"webBlog/model"
)

type MainController struct {
	beego.Controller
}

type TagCode struct {
	Code int
	Msg  string
	Data interface{}
}

var tbNameRichEdit string

func init() {
	tbNameRichEdit = beego.AppConfig.String("tb.richEdit")
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) AddBlog() {
	c.TplName = "addBlog.html"
}

func (c *MainController) UEditorDemo() {
	c.Data["LocalIPPort"] = beego.AppConfig.String("LocalIPPort")
	c.TplName = "ueditorDemo.html"
}

func (c *MainController) SaveRichEdit() {
	var tagCode TagCode
	var richEdit model.RichEdit
	richEdit.Content = "人生何处不相逢，相逢何必曾相识"
	o := orm.NewOrm()
	qs := o.QueryTable(tbNameRichEdit)
	err := qs.Filter("id", 1).One(&richEdit)
	if err!=nil{
		tagCode.Code = 1
		tagCode.Msg = err.Error()
		c.Data["json"] = tagCode
		c.ServeJSON()
		return
	}
	if _, err := o.Update(&richEdit,"Content"); err != nil {
		fmt.Println("更新数据失败！",err.Error())
		tagCode.Code = 2
		tagCode.Msg = err.Error()
		c.Data["json"] = tagCode
		return
	}
	tagCode.Code = 0
	tagCode.Msg = "success"
	c.Data["json"] = tagCode
	c.ServeJSON()
}