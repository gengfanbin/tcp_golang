package base

import (
	"tcp_service/config"
	"tcp_service/hook"
	"tcp_service/pool"
	"tcp_service/utils"
)

//  分发消息 msg 消息体 socket_id 客户端id
func Router(msg []byte, socket_id int) {
	router_id, msg_params, socket_id, pass := MessageBefore(msg, socket_id)
	if !pass {
		return
	}

	// Server 容错处理
	defer func() {
		panic_err := recover()
		if panic_err != nil {
			config.Log("SERVER ERROR: \r\n\t", router_id, ":", panic_err)
		}
	}()

	config.Router(router_id, msg_params, socket_id)

	hook.MessageAfterHook(router_id, msg_params, socket_id)

}

// 消息解析
func MessageBefore(msg []byte, socket_id int) (interface{}, interface{}, int, bool) {

	// 对消息按照指定格式解析
	router_id, msg_params, msg_status := utils.MsgAnalysis(msg)

	pass := true
	if !msg_status {
		cli_addr := pool.SocketPool.UserItem[socket_id].Conn.RemoteAddr().String()
		config.Log("ROUTER ERROR: \r\n\t客户端地址: " + cli_addr + "\r\n\terror msg: " + "消息解析失败")
		pass = false
	} else {
		router_id, msg_params, socket_id, pass = hook.MessageBeforeHook(router_id, msg_params, socket_id)
	}

	return router_id, msg_params, socket_id, pass
}
