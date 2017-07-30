package main

import (
	_ "gisbo/routers"
	"strings"

	"gisbo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	//orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", "root:jkflwurn@tcp(127.0.0.1:3307)/gisbo")

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)
	jdbc := []string{
		beego.AppConfig.String("jdbc.username"),
		":",
		beego.AppConfig.String("jdbc.password"),
		"@tcp(",
		beego.AppConfig.String("jdbc.host"),
		":",
		beego.AppConfig.String("jdbc.port"),
		")/",
		beego.AppConfig.String("jdbc.dbname")}
	connectionString := strings.Join(jdbc, "")
	//DebugLog(connectionString)
	orm.RegisterDataBase("default", "mysql", connectionString, 30)
	orm.RegisterModel(
		new(models.User))
	//new(models.Topic),
	//new(models.Section),
	//new(models.Reply),
	//new(models.ReplyUpLog),
	//new(models.Role),
	//new(models.Permission)

	//orm.RunSyncdb("default", false, true)
	orm.Debug = true
	//models.Init()
}
func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
}
