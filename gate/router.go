package gate

import (
	//"qpgame/game"
	"qpgame/dbcenter"
	"qpgame/login"
	"qpgame/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	//msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.ClientMsg{}, login.ChanRPC)

	//msg.Processor.SetRouter(&msg.PlatFormMsg{}, login.ChanRPC)

	//msg.Processor.SetRouter(&msg.GameMsg{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ReConnect{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.Logout{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.Heatbeat{}, login.ChanRPC)

	//room msg
	msg.Processor.SetRouter(&msg.RoomMsg{}, dbcenter.ChanRPC)
}
