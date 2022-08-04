package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

/*
* 功能：消息编码
* 参数：string 消息内容
 */
func TcpBufEncode(message string) ([]byte, error) {
	// 读取消息的长度转换成uint32类型（4字节）
	var length = uint32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入包体
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

/*
* 功能：消息解码
* 参数：消息读取缓冲区
 */
func TcpBufDecode(reader *bufio.Reader) ([]byte, error) {
	// 读消息长度
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length uint32
	err := binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}
	// buffer返回缓冲中现有的可读的字节数
	if uint32(reader.Buffered()) < length+4 {
		return nil, err
	}
	// 读取真正的数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return nil, err
	}
	return pack[4:], nil
}
