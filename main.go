package main

import (
	_ "dao-service/bussiness-dao-service/routers"

	"model"

	"github.com/astaxie/beego"
)

func main() {
	var cfg = beego.AppConfig
	dbname := cfg.String("dbname")
	dbdriver := cfg.String("dbdriver")
	dbaccount := cfg.String("dbaccount")
	dbaddr := cfg.String("dbaddr")
	dbconnum, _ := cfg.Int64("dbconnum")

	err := model.InitEnv(dbname, dbdriver, dbaccount, dbaddr, dbconnum)
	if err != nil {
		beego.Debug(err)
		return
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
