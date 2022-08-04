package server

import (
	"tcp_service/pool"
)

// 全局聊天
func ChatOfAll(msg_params interface{}, socket_id int) {
	params := msg_params.(map[string]interface{})
	if params["message"] != nil {
		users := make([]int, 0)
		for _, user := range pool.SocketPool.UserItem {
			users = append(users, user.Id)
		}
		pool.SendMessage(params["message"].(string), users)
	}
}

// 房间聊天
func ChatOfRoom(msg_params interface{}, socket_id int) {
	params := msg_params.(map[string]interface{})
	if params["room_id"] != nil && params["message"] != nil {
		room_id := params["room_id"].(int)
		room := Rooms[room_id]
		users := make([]int, 0)
		for _, user := range room.Room_user {
			users = append(users, user.Id)
		}
		pool.SendMessage(params["message"].(string), users)
	}
}
