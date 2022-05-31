package main

import (
	"fmt"
	_ "fyouku/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defaultDb, _ := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	err := orm.RegisterDataBase("default", "mysql", defaultDb, orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	if err != nil {
		fmt.Println("注册数据库错误：", err)
		return
	}

	beego.Run()
}
