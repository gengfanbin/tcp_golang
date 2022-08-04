package base

import (
	"bufio"
	"net"
	"os"
	"strings"
	"tcp_service/config"
	"tcp_service/hook"
	"tcp_service/pool"
	"tcp_service/utils"
)

func TcpInit() {
	config.Log("启动服务")
	ln, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		config.Log("ERROR: 10001 网络端口监听失败, ", err)
		return
	}
	go linsten_tcp_Accept(ln)

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if inputInfo == "close" { // 退出客户端逻辑
			close_status := hook.CloseServerHook()
			if close_status {
				return
			}
		}
	}
}

func linsten_tcp_Accept(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			config.Log("ERROR: 10002 网络连接失败, ", err)
			continue
		}
		if hook.CliInsertBeforeHook(conn) {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	// Socket 容错处理
	defer func() {
		panic_err := recover()
		if panic_err != nil {
			config.Log("SOCKET ERROR: ", panic_err)
		}
	}()

	config.Log("接收到连接", conn.LocalAddr().String())

	socket_id, scoket_msg, status := pool.Add(0, conn)
	if !status {
		config.Log("ERROR: 10003 添加连接池失败: ", "\r\n\t 客户端地址:  ", conn.LocalAddr().String(), "\t错误信息: ", scoket_msg)
		return
	}

	println("当前连接数: ", pool.SocketPool.Size)
	// 处理协议数据
	reader := bufio.NewReader(conn)
	for {
		msg, err := utils.TcpBufDecode(reader)
		if err != nil {
			config.Log("ERROR: 10004 断开链接: ", "\r\n\t 客户端地址:  ", conn.LocalAddr().String(), "\t错误信息: ", err)
			break
		}
		pass := true
		if msg != nil && pass {
			// 分发消息
			Router(msg, socket_id)
		}
	}

	// 关闭连接
	hook.CliCloseHook(socket_id)
}
