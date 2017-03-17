package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// Database
	connectionString   = "benchmarkdbuser:benchmarkdbpass@tcp(TFB-database:3306)/hello_world"
	macIdleConnection  = 30
	maxConnectionCount = 256
)

func init() {
	orm.RegisterModel(new(World), new(Fortune))
	orm.RegisterDataBase("default", "mysql", connectionString, macIdleConnection, maxConnectionCount)
}
