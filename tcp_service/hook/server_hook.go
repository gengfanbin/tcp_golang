package hook

// 关闭服务器 返回一个bool值 标识是否关闭成功
// 返回true 主进程执行退出逻辑
// 返回false 主进程继续执行服务不会关闭
func CloseServerHook() bool {
	println("开始执行关闭服务器")

	return true
}
