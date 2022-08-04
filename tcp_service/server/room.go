package server

import (
	"encoding/json"
	"tcp_service/pool"
)

type Room struct {
	Room_user     []pool.Cli
	Room_id       int
	Room_name     string
	Room_type     int
	Room_password string
}

var Rooms = make(map[int]Room, 0)

// 创建房间
func CreateRoom(msg_params interface{}, socket_id int) {
	params := msg_params.(map[string]interface{})
	user := pool.SocketPool.UserItem[socket_id]
	if params["room_id"] != nil {
		room_id := params["room_id"].(int)
		room_user := make([]pool.Cli, 0)
		room_user = append(room_user, user)
		Rooms[room_id] = Room{
			Room_user:     room_user,
			Room_id:       room_id,
			Room_name:     params["room_name"].(string),
			Room_type:     params["room_type"].(int),
			Room_password: params["Room_password"].(string),
		}
		pool.SendMessage("创建房间成功", []int{socket_id})
	} else {
		pool.SendMessage("创建房间失败", []int{socket_id})
	}
	println("创建房间模块:", Rooms)
}

// 查询房间
func SearchRoom(msg_params interface{}, socket_id int) {
	rooms := make([]Room, 0)
	for _, v := range Rooms {
		rooms = append(rooms, v)
	}
	req_room, err := json.Marshal(rooms)
	if err != nil {
		pool.SendMessage(string(req_room), []int{socket_id})
	}
}

// 加入房间
func InRoom(msg_params interface{}, socket_id int) {
	params := msg_params.(map[string]interface{})
	if params["room_id"] != nil {
		room_id := params["room_id"].(int)
		room := Rooms[room_id]
		room.Room_user = append(room.Room_user, pool.SocketPool.UserItem[socket_id])
		pool.SendMessage("加入房间成功", []int{socket_id})
	} else {
		pool.SendMessage("加入房间失败", []int{socket_id})
	}
}

// 退出房间
func OutRoom(msg_params interface{}, socket_id int) {
	params := msg_params.(map[string]interface{})
	if params["room_id"] != nil {
		room_id := params["room_id"].(int)
		room := Rooms[room_id]
		if len(room.Room_user) < 2 {
			delete(Rooms, room_id)
		} else {
			for i, v := range room.Room_user {
				if v.Id == socket_id {
					room.Room_user = append(room.Room_user[:i], room.Room_user[i+1:]...)
					break
				}
			}
		}
		pool.SendMessage("退出房间成功", []int{socket_id})
	}
}
