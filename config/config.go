package config

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type ConfConnMysql = struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Db   string `json:"db" yaml:"db"`
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
}

var (
	AdminPassword string
	ConnMysql     *ConfConnMysql
	DbDefault     *gorm.DB
)

func Init() {
	var err error

	AdminPassword, err = beego.AppConfig.String("adminPassword")
	if nil != err {
		AdminPassword = ""
	}

	ConnMysql = &ConfConnMysql{}

	ConnMysql.Host, _ = beego.AppConfig.String("connMysql::mysqlHost")
	ConnMysql.Port, _ = beego.AppConfig.Int("connMysql::mysqlPort")
	ConnMysql.Db, _ = beego.AppConfig.String("connMysql::mysqlDb")
	ConnMysql.User, _ = beego.AppConfig.String("connMysql::mysqlUser")
	ConnMysql.Pass, _ = beego.AppConfig.String("connMysql::mysqlPass")
	timeout := "3s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", ConnMysql.User, ConnMysql.Pass, ConnMysql.Host, ConnMysql.Port, ConnMysql.Db, timeout)
	// fmt.Println(dsn)
	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	DbDefault, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DbDefault = DbDefault.Omit(clause.Associations)
	DbDefault = DbDefault.Debug()
	fmt.Println("连接数据库成功")
}

func CheckAdminPassword(pass string) bool {
	return "" != pass && "" != AdminPassword && pass == AdminPassword
}
