package internal

import (
	"fmt"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	. "qpgame/base"
	"qpgame/dbcenter"
	"qpgame/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	token := ""
	adata := a.UserData()
	if adata != nil {
		token = adata.(string)
	}

	if c, ok := Clients[token]; ok {
		nmsg := msg.NormalMsg{}
		nmsg.Code = MSG_DBCENTER_LOGIN_NOTICE
		nmsg.Msg = c.UserID

		c.Agent = a
		c.IsOnline = true
		Clients[token] = c

		dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)
		//dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)
		fmt.Println("rpcNewAgent", a.RemoteAddr().String(), fmt.Sprintf("%d", c.UserID))
	} else {
		fmt.Println("rpcNewAgent", a.RemoteAddr().String())
	}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	token := ""
	adata := a.UserData()
	if adata != nil {
		token = adata.(string)
	}
	//ip_port := a.RemoteAddr().String()
	if c, ok := Clients[token]; ok {
		nmsg := msg.NormalMsg{}
		nmsg.Code = MSG_DB_CENTER_DISCONNECT_NOTICE
		nmsg.Msg = c.UserID

		c.Agent.Destroy()
		c.IsOnline = false
		Clients[token] = c

		dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)
		//err := dbcenter.ChanRPC.Call0("DBCENTER", &nmsg, a)
		//fmt.Println(err)
		log.Debug("user...... rpcCloseAgent...%s, %s, %d", a.RemoteAddr().String(), token, c.UserID)
	} else {
		log.Debug("rpcCloseAgent...%s", a.RemoteAddr().String())
	}
}
