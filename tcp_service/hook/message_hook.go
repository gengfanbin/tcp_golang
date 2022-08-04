package hook

// 处理消息之前的钩子
// 返回四个值
// 第一个值为interface{}类型值 返回处理后的路由id
// 第二个为interface{}类型值 返回处理后的message
// 第三个为int值 返回处理后的socket_id
// 第四个为bool值 标识是否接入消息
func MessageBeforeHook(router_id interface{}, msg_params interface{}, socket_id int) (interface{}, interface{}, int, bool) {

	return router_id, msg_params, socket_id, true
}

// 处理消息之后的钩子 无返回值
func MessageAfterHook(router_id interface{}, msg_params interface{}, socket_id int) {

}
