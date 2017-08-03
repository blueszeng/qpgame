package internal

import (
	. "qpgame/base"
	. "qpgame/base/function"
	"qpgame/base/model"
	"qpgame/msg"
	//"qpgame/util"
	"reflect"
	//"strconv"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.RoomMsg{}, handlerRoomMsg)

	skeleton.RegisterChanRPC("DBCENTER", onMessage)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func onMessage(args []interface{}) {
	m := args[0].(*msg.NormalMsg)
	code := m.Code
	log.Debug("DBCENTER onMessage : %v", code)
	switch code {
	case MSG_DBCENTER_LOGIN_NOTICE:
		onLogin(args)
	case MSG_DBCENTER_LOGOUT_NOTICE:
		onDsiconnect(args)
	case MSG_DB_CENTER_DISCONNECT_NOTICE:
		onDsiconnect(args)
	default:
		//fmt.Printf("Default")
	}

}

func onLogin(args []interface{}) {
	nmsg := msg.ClientMsg{}
	m := args[0].(*msg.NormalMsg)
	a := args[1].(gate.Agent)
	user_id := m.Msg.(uint)

	log.Debug("onLogin %v", a.UserData().(string))

	u, ok := g_user_mgr.Login(user_id, a)
	if ok {
		nmsg.Code = MSG_CLIENT_USERINFO_RSP

		nmsg.Msg = g_user_mgr.GetUserBaseInfo(u)
		SendMsg(a, &nmsg)

		//proc room info
		if u.RoomID > 0 {
			g_roomMgr.ReJoinRoom(u)
			g_roomMgr.SendRoomBaseInfo(u)

			//boadcast to same
			g_roomMgr.BoadcastRoomsInfo(u.RoomID, u.ID, MSG_ROOM_STATUS_CHANGE_NOTICE)
		}
	} else {
		//异常的时候会没有用户信息
	}
}

func onDsiconnect(args []interface{}) {
	//nmsg := msg.NormalMsg{}
	m := args[0].(*msg.NormalMsg)
	//a := args[1].(gate.Agent)

	g_user_mgr.Logout(m.Msg.(uint))

	u := g_user_mgr.Get(m.Msg.(uint))
	//proc room info
	if u != nil && u.RoomID > 0 {
		g_roomMgr.ReJoinRoom(u)
		//g_roomMgr.ProcUserRoom(u)

		//boadcast to same
		g_roomMgr.BoadcastRoomsInfo(u.RoomID, u.ID, MSG_ROOM_STATUS_CHANGE_NOTICE)
	}
}

func handlerRoomMsg(args []interface{}) {
	m := args[0].(*msg.RoomMsg)
	code := m.Code
	log.Debug("handlerRoomMsg onMessage : %v", code)

	switch code {
	case MSG_ROOM_CREATE_REQ:
		onCreateRoom(args)
	case MSG_ROOM_JOIN_REQ:
		onJoinRoom(args)
	case MSG_ROOM_APPLY_DISSOLVE_REQ:
		onDissolveRoom(args)
	case MSG_ROOM_AGREE_DISSOLVE_REQ:
		onAnswerDissolveRoom(args)
	case MSG_ROOM_LEAVE_REQ:
		onLeaveRoom(args)
	default:
		//fmt.Printf("Default")
	}
}

//token check
func _checkToken(a gate.Agent) (*model.User, bool) {
	nmsg := msg.ClientMsg{}

	token := a.UserData().(string)
	u, ok := g_user_mgr.GetUserByToken(token)
	if !ok {
		//erro proc
		nmsg.Code = MSG_CLIENT_TOKEN_EXPIRED_NOTICE
		nmsg.Status = SUM_MSG_STATUS_ERROR
		nmsg.Msg = "token expired"
		SendMsg(a, &nmsg)

		log.Debug("onCreateRoom get user info wrong token=%s", token)
		return nil, false
	}
	return u, true
}

func onCreateRoom(args []interface{}) {
	//check have card to n creat room

	nmsg := msg.ClientMsg{}
	//m := args[0].(*msg.RoomMsg)
	a := args[1].(gate.Agent)

	u, ok := _checkToken(a)
	if !ok {
		return
	}

	game_conf := msg.GameConf{}
	game_conf.ID = 1
	game_conf.Name = "kwx"
	game_conf.PeopleNumbers = 3
	game_conf.GameNumbers = 8
	room, status := g_roomMgr.Create(u, game_conf)
	if status == SUM_MSG_STATUS {
		nmsg.Code = MSG_ROOM_CREATE_RSP
		nmsg.Status = status
		nmsg.Msg = room.GetRoomBaseInfo();
		SendMsg(a, &nmsg)
		return
	} else {
		//异常的时候会没有用户信息
		nmsg.Code = MSG_ROOM_CREATE_RSP
		nmsg.Status = status
		nmsg.Msg = ""
		SendMsg(a, &nmsg)
		return
	}
}

func onJoinRoom(args []interface{}) {
	nmsg := msg.ClientMsg{}
	m := args[0].(*msg.RoomMsg)
	a := args[1].(gate.Agent)
	room_id := m.RoomID

	u, ok := _checkToken(a)
	if !ok {
		return
	}
	room, status := g_roomMgr.Join(room_id, u)
	if status == SUM_MSG_STATUS {
		nmsg.Code = MSG_ROOM_JOIN_RSP
		nmsg.Status = status
		nmsg.Msg = room.GetRoomBaseInfo();
		SendMsg(a, &nmsg)
		log.Debug("onJoinRoom userid=%d, room_id=%d ", u.ID, room_id)
		g_roomMgr.BoadcastRoomsInfo(room_id, u.ID, MSG_ROOM_SYS_INFO_NOTICE)
		room.SetReady(u)
	} else {
		nmsg.Code = MSG_ROOM_JOIN_RSP
		nmsg.Status = status
		nmsg.Msg = ""
		SendMsg(a, &nmsg)
		return
	}
}

func onDissolveRoom(args []interface{}) {
	nmsg := msg.NormalMsg{}
	m := args[0].(*msg.RoomMsg)
	a := args[1].(gate.Agent)
	room_id := m.RoomID

	u, ok := _checkToken(a)
	if !ok {
		return
	}
	_, status := g_roomMgr.Dissolve(room_id, u)

	nmsg.Code = MSG_ROOM_AGREE_DISSOLVE_RSP
	SendMsg(a, &nmsg)
	g_roomMgr.BoadcastRoomsInfo(room_id, u.ID, MSG_ROOM_SYS_DISSOLVE_NOTICE)
	log.Debug("onDissolveRoom done userid=%d, room_id=%d, status=%d ", u.ID, room_id, status)

}

func onAnswerDissolveRoom(args []interface{}) {
	nmsg := msg.ClientMsg{}
	m := args[0].(*msg.RoomMsg)
	a := args[1].(gate.Agent)
	room_id := m.RoomID

	is_agree := m.Msg.(int)
	u, ok := _checkToken(a)
	if !ok {
		return
	}
	g_roomMgr.AnswerDissolve(room_id, is_agree, u)

	nmsg.Code = MSG_ROOM_AGREE_DISSOLVE_RSP
	SendMsg(a, &nmsg)

	g_roomMgr.BoadcastRoomsInfo(room_id, u.ID, MSG_ROOM_SYS_DISSOLVE_NOTICE)
}

func onLeaveRoom(args []interface{}) {
	nmsg := msg.ClientMsg{}
	m := args[0].(*msg.RoomMsg)
	a := args[1].(gate.Agent)
	room_id := m.RoomID

	u, ok := _checkToken(a)
	if !ok {
		return
	}
	if u.RoomID != room_id {
		//do something
		nmsg.Code = MSG_ROOM_LEAVE_RSP
		nmsg.Status = ROOM_STATUS_NOT_EXIST
		log.Debug("onLeaveRoom userid=%d, room_id=%d==u.room_id=%d ", u.ID, room_id, u.RoomID)
		
	} else {
		_, status := g_roomMgr.Leave(room_id, u)
		nmsg.Code = MSG_ROOM_LEAVE_RSP
		nmsg.Status = status
	}
	SendMsg(a, &nmsg)
	g_roomMgr.BoadcastRoomsInfo(room_id, 0, MSG_ROOM_SYS_INFO_NOTICE)
}
