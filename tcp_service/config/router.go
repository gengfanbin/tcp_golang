package config

import (
	"tcp_service/pool"
	"tcp_service/server"
)

//  分发消息 socket_id 客户端ID msg_params 消息体 router_id路由ID
func Router(router_id interface{}, msg_params interface{}, socket_id int) {
	switch router_id.(string) {
	case "10011": // 创建房间
		server.CreateRoom(msg_params, socket_id)
	case "10012": // 创建房间
		server.SearchRoom(msg_params, socket_id)
	case "10013": // 加入房间
		server.InRoom(msg_params, socket_id)
	case "10014": // 离开房间
		server.OutRoom(msg_params, socket_id)
	case "10021": // 发送消息
		server.ChatOfAll(msg_params, socket_id)
	case "10022": // 发送消息
		server.ChatOfRoom(msg_params, socket_id)
	default:
		Log("路由服务匹配失败: ", "\r\n\t客户端地址: ", pool.SocketPool.UserItem[socket_id].Conn.RemoteAddr().String(), "\r\n\t路由id: ", router_id)
	}

}
