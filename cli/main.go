package main

import (
	"cli/utils"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Message struct {
	Id     string
	Params map[string]interface{}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12001")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	// 创建房间
	msg := Message{}
	msg.Id = "10011"
	msg.Params = make(map[string]interface{}, 0)
	msg.Params["room_id"] = "1"

	msg_json, err := json.Marshal(msg)
	println("msg_json : ", string(msg_json), "||", err != nil)

	if err != nil {
		fmt.Println("json marshal failed, err", err)
		return
	}
	data, err := utils.TcpBufEncode(string(msg_json))
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return
	}
	conn.Write(data)

	time.Sleep(time.Duration(10) * time.Second)
}
