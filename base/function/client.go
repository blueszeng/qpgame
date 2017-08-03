package function

import (
	//. "qpgame/base"
	//"qpgame/game"
	//"qpgame/msg"

	"github.com/name5566/leaf/gate"
)

//easy to logger message
func SendMsg(a interface{}, m interface{}) {
	agent := a.(gate.Agent)
	agent.WriteMsg(m)

}
