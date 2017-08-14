package internal

import (
	"fmt"
	"time"

	"reflect"
	//"strconv"

	//"qpgame/game"
	. "qpgame/base"
	. "qpgame/base/function"

	"qpgame/base/model"
	"qpgame/dbcenter"
	"qpgame/msg"
	. "qpgame/util"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//handler(&msg.ClientMsg{}, onClientMsg)

	handler(&msg.Login{}, handleLogin)
	handler(&msg.ReConnect{}, handleReConnect)
	handler(&msg.Logout{}, handleLogout)
	handler(&msg.Heatbeat{}, handleHeatbeat)
	//go heatbeat()
}

func isRepeatLanding(uid uint) (client, bool) {
	for _, v := range Clients {
		if v.UserID == uid {
			if v.IsOnline {
				return v, true
			}
		}
	}

	return client{}, false
}

/*
func onClientMsg(args []interface{}) {
	//recive msg
	m := args[0].(*msg.ClientMsg)
	fmt.Println("message code:", m.Code)
	switch m.Code {
	case MSG_CLIENT_LOGIN_REQ:
		handleLogin(args)
	case MSG_CLIENT_RECONNECT_REQ:
		onReConnect(args)
	}
}
*/

func handleLogin(args []interface{}) {
	var (
		nmsg       msg.NormalMsg //notice dbcenter
		client_msg msg.ClientMsg //notice client
	)

	//contents := args[0].(*msg.ClientMsg).Msg(map[string]interface{})
	//recive msg
	m := args[0].(*msg.Login)
	//client agent
	a := args[1].(gate.Agent)

	user := model.User{}
	param := make(map[string]string)
	param["user_name = ?"] = m.Name
	user.GetOne(param)
	fmt.Println(user.Password, user.ID)

	if user.ID > 0 && CheckPassword(m.Password, user.Password) {
		//if repeat landing
		if c, ok := isRepeatLanding(user.ID); ok {
			log.Debug("repeat landing %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())

			//notice old user repeat landing
			client_msg.Msg = a.RemoteAddr().String()
			client_msg.Status = SUM_MSG_STATUS_SUCCESS
			client_msg.Code = MSG_CLIENT_REPEAT_NOTICE
			SendMsg(c.Agent, &client_msg)

			//response
			client_msg.Msg = ""
			client_msg.Status = SUM_MSG_STATUS_REPEAT
			client_msg.Code = MSG_CLIENT_LOGIN_RSP
			SendMsg(a, &client_msg)

		} else {
			log.Debug("login success %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())

			token := GetMd5String(a.RemoteAddr().String() + time.Local.String() + fmt.Sprintf("%v", user.ID))

			a.SetUserData(token)

			c := client{}
			c.Token = token
			c.UserID = user.ID
			c.Agent = a
			c.IsOnline = true
			Clients[token] = c

			nmsg.Code = MSG_DBCENTER_LOGIN_NOTICE
			nmsg.Msg = user.ID
			dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)

			client_msg.Msg = token
			client_msg.Code = MSG_CLIENT_LOGIN_RSP
			SendMsg(a, &client_msg)

		}
	} else {
		log.Debug("登陆失败 %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())
		//如何不存在就加入
		client_msg.Code = MSG_CLIENT_LOGIN_RSP
		client_msg.Status = SUM_MSG_STATUS_ERROR
		client_msg.Msg = "登陆密码或者帐号错误"

		SendMsg(a, &client_msg)
	}
}

//recontect
func handleReConnect(args []interface{}) {
	var (
		nmsg       msg.NormalMsg //notice dbcenter
		client_msg msg.ClientMsg //notice client
	)

	//recive msg
	m := args[0].(*msg.ReConnect)
	token := m.Token

	//client agent
	a := args[1].(gate.Agent)
	client, ok := Clients[token]

	if ok {
		//if not the same ip:port reconnect
		if client.IsOnline && client.Agent.RemoteAddr().String() != a.RemoteAddr().String() {
			log.Debug("repeat reconnect success %v/%v@%v", client.UserID, client.Token)
			client_msg.Msg = a.RemoteAddr().String()
			client_msg.Status = SUM_MSG_STATUS_SUCCESS
			client_msg.Code = MSG_CLIENT_REPEAT_NOTICE
			client.Agent.WriteMsg(&client_msg)

			//client.Agent.Destroy()
			//response
			client_msg.Msg = ""
			client_msg.Status = SUM_MSG_STATUS_REPEAT
			client_msg.Code = MSG_CLIENT_LOGIN_RSP
		} else {
			log.Debug("reconnect success %v/%v@%v", client.UserID, client.Token)
			a.SetUserData(token)
			client.Agent = a
			client.IsOnline = true
			Clients[token] = client

			client_msg.Msg = token
			client_msg.Code = MSG_CLIENT_RECONNECT_RSP
			client_msg.Status = 0

			nmsg.Code = MSG_DBCENTER_LOGIN_NOTICE
			nmsg.Msg = client.UserID
			dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)
		}
		//err := dbcenter.ChanRPC.Call0("DBCENTER", &nmsg, a)
		//fmt.Println(err)

	} else {
		log.Debug("reconnect falied %v", token)
		//如何不存在就加入
		client_msg.Code = MSG_CLIENT_RECONNECT_RSP
		client_msg.Status = SUM_MSG_STATUS_ERROR
		client_msg.Msg = "token过期"
	}

	SendMsg(a, &client_msg)
	//通知游戏服务器此用户登陆成功
	//game.ChanRPC.Go("NewAgent", base.MSG_LOGIN, m, a)
}

func handleLogout(args []interface{}) {
	var (
		nmsg msg.NormalMsg //notice dbcenter
		//client_msg msg.NormalMsg //notice client
	)

	//recive msg
	m := args[0].(*msg.Logout)
	token := m.Token

	//client agent
	a := args[1].(gate.Agent)

	client, ok := Clients[token]

	if ok {
		log.Debug("logout success %v/%v@%v", client.UserID, client.Token)
		delete(Clients, token)

		nmsg.Code = MSG_DBCENTER_LOGOUT_NOTICE
		nmsg.Msg = client.UserID
		dbcenter.ChanRPC.Go("DBCENTER", &nmsg, a)
		//err := dbcenter.ChanRPC.Call0("DBCENTER", &nmsg, a)
		//fmt.Println(err)
		//client_msg.Msg = token
		//client_msg.Code = MSG_CLIENT_RECONNECT_RSP
		//client_msg.Status = 0
	} else {
		/*log.Debug("reconnect falied %v", token)
		//如何不存在就加入
		client_msg.Code = MSG_CLIENT_RECONNECT_RSP
		client_msg.Status = 1
		client_msg.Msg = "token过期"*/
	}

	//SendMsg(a, &client_msg)
	//通知游戏服务器此用户登陆成功
	//game.ChanRPC.Go("NewAgent", base.MSG_LOGIN, m, a)
}

func handleHeatbeat(args []interface{}) {
	m := args[0].(*msg.Heatbeat)
	nmmsg := msg.ClientMsg{}
	//client agent
	a := args[1].(gate.Agent)
	token := ""
	adata := a.UserData()
	if adata != nil {
		token = adata.(string)
	}

	nmmsg.Code = 0

	if c, ok := Clients[token]; ok {
		nmmsg.Msg = time.Local.String() + "," + m.HB
		SendMsg(c.Agent, &nmmsg)
	} else {
		nmmsg.Msg = time.Local.String() + "," + m.HB
		SendMsg(a, &nmmsg)
	}

	log.Debug("heartbeat .... %v", m.HB)

}

//heartbeat
func heatbeat() {
	c := time.Tick(60 * time.Second)
	for _ = range c {
		//fmt.Printf("%v \n", now)
		for _, v := range Clients {
			if v.IsOnline {
				nmmsg := msg.NormalMsg{}
				nmmsg.Code = 0
				nmmsg.Msg = v.UserID
				log.Debug("heatbeat:%v", v.UserID)
				SendMsg(v.Agent, &nmmsg)
			}
		}
	}
}
