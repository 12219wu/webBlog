package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RichEdit struct {
	Id int64    `orm:"column(id)`
	Content string  `orm:"column(content)`
}

func (tb *RichEdit) TableName() string {
	return beego.AppConfig.String("tb.richEdit")
}

func init() {
	fmt.Println("MySQL init")
	
	dbUser := beego.AppConfig.String("db.User")
	dbPwd := beego.AppConfig.String("db.Pwd")
	dbHost := beego.AppConfig.String("db.Host")
	dbPort := beego.AppConfig.String("db.Port")
	dbName := beego.AppConfig.String("db.Name")
	conn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	fmt.Println(conn)
	// set default database
	orm.RegisterDataBase("default", "mysql", conn, 30)
	// wyd
	orm.RegisterModel(new(RichEdit))
	// create table
	orm.RunSyncdb("default", false, true)
}