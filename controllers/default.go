package controllers

import (
	"goweb/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "user/index.html"
}

func (c *MainController) User() {
	u := models.User{Id: "golwei"}
	u.ReadDB()
	c.Data["json"] = u
	c.ServeJSON()
}
