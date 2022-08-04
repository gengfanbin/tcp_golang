package pool

import (
	"tcp_service/utils"
)

// 广播消息
func SendMessage(msg string, cli_id []int) {
	for _, c_id := range cli_id {
		for _, user := range SocketPool.UserItem {
			if user.Id == c_id {
				message, err := utils.TcpBufEncode(msg)
				if err != nil {
					user.Conn.Write(message)
				}
			}
		}
	}
}
