package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//	"goweb/models/class"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)                                     //驱动
	orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file")) //连接
	orm.RegisterModel(new(User))                                                   //数据模型
	orm.RunSyncdb("default", false, true)                                          //建立表
}

type User struct {
	Id   string `orm:"pk"`
	Name string
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return
}

func (u User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	return
}
