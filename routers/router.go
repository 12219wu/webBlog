package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"webBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	
    beego.Router("/", &controllers.MainController{})
	beego.Router("/addBlog", &controllers.MainController{},"get:AddBlog")
	beego.Router("/ueditorDemo", &controllers.MainController{},"get:UEditorDemo")
	beego.Router("/richEdit", &controllers.MainController{},"get:SaveRichEdit")
}
