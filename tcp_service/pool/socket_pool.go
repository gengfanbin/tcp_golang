package pool

import (
	"net"
	"time"
)

type Cli struct {
	Id        int
	StartTime int
	EndTime   int
	Info      map[string]interface{}
	Conn      net.Conn
}

type Pool struct {
	Size     int
	UserItem map[int]Cli
}

// 定义连接池
var SocketPool = InitPoll()

func InitPoll() Pool {
	pool := Pool{
		Size:     0,
		UserItem: make(map[int]Cli, 0),
	}
	return pool
}

// 返回id 错误信息 和 状态
func Add(id int, conn net.Conn) (int, string, bool) {
	// id 必须是唯一的 如果id为0 则自动生成id
	if id == 0 {
		id = int(time.Now().UnixMilli())
	}
	// 如果id已经存在 则返回
	if _, ok := SocketPool.UserItem[id]; ok {
		return 0, "id已经存在", false
	}

	StartTime := int(time.Now().UnixMilli())
	newItem := Cli{
		Id:        id,
		Info:      make(map[string]interface{}, 0),
		StartTime: StartTime,
		EndTime:   StartTime,
		Conn:      conn,
	}
	SocketPool.UserItem[id] = newItem
	SocketPool.Size = len(SocketPool.UserItem)
	return id, "成功", true
}

// 删除链接池中指定ID
func Del(id int) {
	SocketPool.UserItem[id].Conn.Close()
	delete(SocketPool.UserItem, id)
	SocketPool.Size = len(SocketPool.UserItem)
}

// 删除所有链接
func DelAll() {
	for key := range SocketPool.UserItem {
		Del(key)
	}
}
