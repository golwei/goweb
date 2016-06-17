package main

import (
	_ "goweb/routers"

	_ "goweb/models"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

//models models.class
//连接数据库 ->models/db.go
//数据结构体 ->models/class/user.go
//控制器 controllers/user.go
