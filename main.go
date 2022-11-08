package main

import (
	_ "cookie-shop-api/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlConn,err := beego.AppConfig.String("sqlconn")
	if err != nil {
		//beeLogger.Log.Fatal("%s", err)
	}
	orm.RegisterDataBase("default", "mysql", sqlConn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.RouterCaseSensitive = true
	tree := beego.PrintTree()
	methods := tree["Data"].(beego.M)
	for k, v := range methods {
		fmt.Printf("%s => %v\n", k, v)
	}
	beego.Run()
}
