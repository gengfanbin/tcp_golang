package mysql

import (
	"tcp_service/base"
)

// 配置mysql数据库，可以配置多个实例，每项配置为一个数据库实例
var Mysql_test = base.Mysql{
	Db:        nil,
	USERNAME:  "root",
	PASSWORD:  "root",
	NETWORK:   "tcp",
	SERVER:    "127.0.0.1",
	PORT:      "3306",
	DATABASE:  "test",
	CHARSET:   "utf8mb4",
	COLLATION: "utf8mb4_unicode_ci",
}
