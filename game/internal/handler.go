package internal

import (
	//"qpgame/msg"
	"reflect"

	//"github.com/name5566/leaf/gate"
	//"github.com/name5566/leaf/log"
)

func init() {
	//handler(&msg.Hello{}, handleHello)
	//handler(&msg.RoomMsg{}, handlerRoomMsg)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	/*// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(`adsf`)*/
}
