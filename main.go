package main

import (
	_ "fyouku/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defaultDb, _ := config.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultDb)

	beego.Run()
}
