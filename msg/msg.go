package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

//var Processor network.Processor
var Processor = json.NewProcessor()
var ProcessorPB = protobuf.NewProcessor()

//var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&ClientMsg{})

	Processor.Register(&Login{})
	Processor.Register(&ReConnect{})
	Processor.Register(&Logout{})
	Processor.Register(&Heatbeat{})

	//room msg
	Processor.Register(&RoomMsg{})

	Processor.Register(&NormalMsg{})
	Processor.Register(&GameMsg{})
	Processor.Register(&PlatFormMsg{})

	ProcessorPB.Register(&AddressBook{})
}

//用户登陆消息体
type Login struct {
	Name     string
	Password string
}

type ReConnect struct {
	Token string
}

type Logout struct {
	Token string
}

type Heatbeat struct {
	HB string
}

type RoomMsg struct {
	Code   int
	RoomID int
	Msg    interface{}
	//something
}

//普通消息通知
type NormalMsg struct {
	Code   int
	Status int //0-success 1-failed
	Msg    interface{}
}

//recv from client game msg
type GameMsg struct {
	Token string
	Code  int
	Msg   interface{}
}

//recv from client PlatForm msg
type PlatFormMsg struct {
	Token string
	Code  int
	Msg   interface{}
}

//game init msg
type GameConf struct {
	ID            int
	Name          string
	PeopleNumbers int
	GameNumbers   int
}
