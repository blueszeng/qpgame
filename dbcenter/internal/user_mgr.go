package internal

import (
	"fmt"
	. "qpgame/base/function"
	"qpgame/base/model"
	"qpgame/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var (
	g_user_mgr       *UserMgr
	g_tokenAndUidMap map[string]uint
)

func init() {
	g_user_mgr = new(UserMgr)
	g_user_mgr.users = make(map[uint]*model.User, 0)
	g_tokenAndUidMap = make(map[string]uint)

}

type UserMgr struct {
	users map[uint]*model.User
}

func (mgr *UserMgr) Get(uid uint) *model.User {
	u, ok := mgr.users[uid]
	if ok {
		return u
	}
	return nil
}
func (mgr *UserMgr) Login(user_id uint, a gate.Agent) (*model.User, bool) {
	u, ok := mgr.users[user_id]
	if ok {
		token := a.UserData().(string)

		u.IsOnline = 1
		u.Agent = a
		//fmt.Println(u)
		mgr.users[user_id] = u

		g_tokenAndUidMap[token] = user_id
		log.Debug("MEM=>userid=%d, name=%s, isonline=%d, room_id=%d", u.ID, u.UserName, u.IsOnline, u.RoomID)

	} else {
		//reload user
		new_user := new(model.User)
		param := make(map[string]string)
		param["id = ?"] = fmt.Sprintf("%d", user_id)
		new_user.GetOne(param)
		if new_user.ID == 0 {
			fmt.Println("reload failed", user_id)

			return new_user, false
		}
		new_user.IsOnline = 1
		new_user.Agent = a
		u = new_user

		token := a.UserData().(string)
		g_tokenAndUidMap[token] = user_id

		log.Debug("DB=>userid=%d, name=%s, isonline=%d, room_id=%d", u.ID, u.UserName, u.IsOnline, u.RoomID)
	}
	mgr.users[user_id] = u
	return u, true
}

func (mgr *UserMgr) Logout(user_id uint) bool {
	u, ok := mgr.users[user_id]
	if ok {
		u.IsOnline = 0

		//nmsg.Code = MSG_CLIENT_USERINFO
		//nmsg.Msg = u
		token := u.Agent.UserData().(string)
		delete(g_tokenAndUidMap, token)

		log.Debug("MEM=>userid=%d, name=%s, isonline=%d, room_id=%d", u.ID, u.UserName, u.IsOnline, u.RoomID)
		mgr.users[user_id] = u

		return true
	} else {
		//reload user
		return false
	}
}

//return roomid or reutnr 0
func (mgr *UserMgr) IsInRoom(uid uint) int {
	u, ok := mgr.users[uid]
	if ok && u.RoomID != 0 {
		log.Debug("IsInRoom user:%d is in room:%d", uid, u.RoomID)
		return u.RoomID
	} else {
		log.Debug("IsInRoom user:%d is not in room", uid)
		return 0
	}
}

func (mgr *UserMgr) GetUserByToken(token string) (*model.User, bool) {
	uid, ok := g_tokenAndUidMap[token]
	if ok {
		u, ok := mgr.users[uid]
		if ok {
			return u, true
		}
	}

	return nil, false
}

func (mgr *UserMgr) SendMsg(uid uint, m interface{}) {
	u, ok := mgr.users[uid]
	if ok {
		SendMsg(u.Agent, m)
	}
}

func (rmg *UserMgr) GetUserBaseInfo(u *model.User) *msg.UserInfo{
		user_info := new(msg.UserInfo)
		user_info.ID = u.ID
		user_info.UserName = u.UserName
		user_info.IsOnline = u.IsOnline
		user_info.RoomID = u.RoomID
		user_info.RoomCards = u.RoomCards
		user_info.Status = u.Status
		user_info.Nickname = u.Nickname
		user_info.Type = u.Type

		return user_info
}
