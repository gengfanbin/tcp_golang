package main

import (
	"tcp_service/base"
	"tcp_service/config"
	"tcp_service/mysql"
)

func main() {
	mysql.Mysql_test.DBInit()
	config.LogInit()
	base.TcpInit()
}
