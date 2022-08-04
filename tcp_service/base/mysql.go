package base

// mysql封装不要直接修改这个文件
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Db        *sql.DB
	USERNAME  string
	PASSWORD  string
	NETWORK   string
	SERVER    string
	PORT      string
	DATABASE  string
	CHARSET   string
	COLLATION string
}

func (mysql *Mysql) DBInit() {
	println("启动SQL服务: ", mysql.SERVER, ":", mysql.PORT, "/", mysql.DATABASE)

	dsn := mysql.USERNAME +
		":" + mysql.PASSWORD +
		"@" + mysql.NETWORK +
		"(" + mysql.SERVER +
		":" + mysql.PORT +
		")/" + mysql.DATABASE +
		"?charset=" + mysql.CHARSET +
		"&collation=" + mysql.COLLATION +
		"&parseTime=1&multiStatements=1"

	// open函数只是验证格式是否正确，并不是创建数据库连接
	var err error
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("数据库连接参数不正确")
	} else {
		mysql.Db = db
	}
}
