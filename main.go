package main

import (
	"fmt"
	_ "sonnn/docs"
	_ "sonnn/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// user:pass@host/dbname?params_extra
	orm.RegisterDataBase("default", "mysql", "sonnn:sonnn1987@/mrapi?charset=utf8")
	//orm.RegisterDataBase("default", "mysql", "root:RMAYziHR@45.117.80.63/golang?charset=utf8")

	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
