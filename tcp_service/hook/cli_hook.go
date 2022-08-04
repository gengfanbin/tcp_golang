package hook

import (
	"net"
	"tcp_service/pool"
)

// 客户端连接钩子 返回一个bool值
// 为true时接入客户端链接,否则抛弃链接
func CliInsertBeforeHook(conn net.Conn) bool {

	return true
}

// 客户端断开链接钩子
func CliCloseHook(socket_id int) {
	println("链接断开: id", socket_id)
	pool.Del(socket_id)
	println("当前连接数: ", pool.SocketPool.Size)
}
