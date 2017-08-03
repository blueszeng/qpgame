package internal

import (
	"qpgame/base"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer

	Clients = make(map[string]client) //client on login infos
)

type Module struct {
	*module.Skeleton
}

type client struct {
	Token    string
	UserID   uint
	Agent    gate.Agent
	IsOnline bool
}

func init() {
	Clients = make(map[string]client)
}
func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
