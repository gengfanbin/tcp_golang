package utils

import (
	"encoding/json"
)

type Message struct {
	Id     string      `json:"id"`
	Params interface{} `json:"params"`
}

func MsgAnalysis(msg []byte) (msg_id interface{}, msg_params interface{}, msg_status bool) {
	var message Message

	msg_status = false
	if err := json.Unmarshal(msg, &message); err != nil {
		return "", "", msg_status
	}

	msg_id = message.Id
	msg_params = message.Params
	msg_status = true
	return
}
