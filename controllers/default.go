package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
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
	c.TplName = "ueditorDemo.html"
}