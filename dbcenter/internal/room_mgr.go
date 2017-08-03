package internal

import (
	. "qpgame/base"
	. "qpgame/base/function"
	"qpgame/base/model"
	"qpgame/msg"
	"time"

	"github.com/name5566/leaf/log"
)

var (
	g_roomMgr *RoomMgr
)

type RoomMgr struct {
	rooms map[int]*Room
}

func init() {
	g_roomMgr = new(RoomMgr)
	g_roomMgr.rooms = make(map[int]*Room, 0)
}

func (rmg *RoomMgr) CreateRoomID() int {
	room_id := RandInt64(100000, 999999)
	_, ok := rmg.Get(room_id)
	if !ok {
		return room_id
	} else {
		return rmg.CreateRoomID()
	}
}

func (rmg *RoomMgr) IsInRoom(uid uint) {
	/*for v,k := rmg.rooms {

	}*/
	//return true
}
func (rmg *RoomMgr) Create(u *model.User, game_conf msg.GameConf) (*Room, int) {
	if u.RoomID != 0 {
		log.Debug("Create User is in room %d", u.RoomID)
		return nil, SUM_MSG_STATUS_REPEAT
	}

	room := new(Room)
	room.ID = rmg.CreateRoomID()

	if _, ok := rmg.rooms[room.ID]; ok {
		log.Debug("room id %d is be used", room.ID)
		return nil, SUM_MSG_STATUS_REPEAT
	}

	room.Name = game_conf.Name
	room.PeopleNumbers = game_conf.PeopleNumbers
	room.GameNumbers = game_conf.GameNumbers
	room.CreatedAt = time.Now()
	room.CreatedBy = u.ID
	room.Seats = make(map[int]*Seat)
	room.Join(u)
	room.ZhuangIndex = 0

	room.SetReady(u)
	var game IGame
	if 1 == game_conf.ID {
		game = new(GameKWX)
	}

	room.Game = game
	rmg.rooms[room.ID] = room

	return room, SUM_MSG_STATUS_SUCCESS

}

func (rmg *RoomMgr) Get(room_id int) (*Room, bool) {
	room, ok := rmg.rooms[room_id]
	if ok {
		return room, true
	}
	return nil, false
}

func (rmg *RoomMgr) Join(room_id int, u *model.User) (*Room, int) {
	room, ok := rmg.rooms[room_id]
	if ok {
		if room_id == u.RoomID && u.RoomID != 0 {
			log.Debug("aleady in room, userid=%d, room_id=%d ", u.ID, room_id)
			return nil, SUM_MSG_STATUS_REPEAT
		} else {
			if len(room.Seats) == room.PeopleNumbers || room.Status == ROOM_STATUS_PLAYING {
				log.Debug("onJoinRoom to much people, userid=%d, room_id=%d ", u.ID, room_id)
				return nil, SUM_MSG_STATUS_TO_MUCH
			}

			room.Join(u)
			//BoadcastRoomsInfo(room_id, u.ID)
			log.Debug("onJoinRoom userid=%d, room_id=%d, room_status=%d ", u.ID, room_id, room.Status)
			return room, SUM_MSG_STATUS_SUCCESS

		}
	} else {
		log.Debug("onJoinRoom room not exist, userid=%d, room_id=%d ", u.ID, room_id)
		return nil, SUM_MSG_STATUS_ERROR
	}
}

func (rmg *RoomMgr) Leave(room_id int, u *model.User) (*Room, int) {
	room, ok := rmg.Get(room_id)
	if !ok {
		return nil, ROOM_STATUS_NOT_EXIST
	} else {
		if room.Status == ROOM_STATUS_WAIT {
			room.Leave(u)
			u.RoomID = 0
			if len(room.Seats) == 0 {
				//room.Status = ROOM_STATUS_OVER
				delete(rmg.rooms, room_id)
				return nil, SUM_MSG_STATUS_SUCCESS
			} else {
				return room, SUM_MSG_STATUS_SUCCESS
			}

		} else if room.Status == ROOM_STATUS_PLAYING {
			rmg.Dissolve(room_id, u)
			return room, ROOM_STATUS_APL_DIS
		} else if room.Status == ROOM_STATUS_APL_DIS {
			rmg.Dissolve(room_id, u)
			return room, ROOM_STATUS_APL_DIS
		} else if room.Status == ROOM_STATUS_OVER {
			return room, SUM_MSG_STATUS_SUCCESS
		}
	}

	return nil, ROOM_STATUS_NOT_EXIST
}

func (rmg *RoomMgr) Dissolve(room_id int, u *model.User) (*Room, int) {
	room, ok := rmg.Get(room_id)
	if !ok {
		return nil, ROOM_STATUS_NOT_EXIST
	} else {
		if room.Status == ROOM_STATUS_WAIT {
			/*room.Leave(u)
			if len(room.Seats) == 0 {
				room.Status = ROOM_STATUS_OVER
				delete(rmg.rooms, room_id)
				return nil, ROOM_STATUS_OVER
			}
			return room, SUM_MSG_STATUS_SUCCESS*/
		} else if room.Status == ROOM_STATUS_PLAYING {
			room.Dissolve(u)
			return room, ROOM_STATUS_APL_DIS
		} else if room.Status == ROOM_STATUS_APL_DIS {
			room.AgreeDissolve(u)
			return room, SUM_MSG_STATUS_REPEAT

		} else if room.Status == ROOM_STATUS_OVER {
			return room, ROOM_STATUS_OVER
		}
		return nil, SUM_MSG_STATUS_SUCCESS
	}
}

func (rmg *RoomMgr) AnswerDissolve(room_id int, is_agree int, u *model.User) (*Room, int) {
	room, ok := rmg.Get(room_id)
	if !ok {
		return nil, ROOM_STATUS_NOT_EXIST
	} else {
		if is_agree == 1 {
			if room.Status == ROOM_STATUS_WAIT {
			} else if room.Status == ROOM_STATUS_PLAYING {

			} else if room.Status == ROOM_STATUS_APL_DIS {
				if room.AgreeDissolve(u) {
					//can dissolve
					room.Status = ROOM_STATUS_DIS_DONE
					return room, ROOM_STATUS_DIS_DONE
				} else {
					return room, ROOM_STATUS_APL_DIS
				}
			} else if room.Status == ROOM_STATUS_OVER {
				return room, ROOM_STATUS_OVER
			}
		} else {

		}

		return nil, SUM_MSG_STATUS_SUCCESS
	}
}

//logger to db
//clear seats info
func (rmg *RoomMgr) Clear(room_id int) {
	room, ok := rmg.Get(room_id)
	if ok {
		log.Debug("Clear room %d", room_id)
		room.Clear()
		delete(rmg.rooms, room_id)
	}
}

func (rmg *RoomMgr) UpdateStatus(room_id int) int {
	room, ok := rmg.Get(room_id)
	if !ok {

	} else {
		if ROOM_STATUS_OVER == room.Status {
			//down data to db
			if len(room.Seats) == 0 {
				room.Status = ROOM_STATUS_OVER
			} else {
				for k, v := range room.Seats {
					//v.UID
					_ = k
					_ = v
				}
			}
		}
	}

	return 0
}

//relogin join room
func (rmg *RoomMgr) ReJoinRoom(u *model.User) {
	room, ok := rmg.Get(u.RoomID)
	if ok {
		for k, v := range room.Seats {
			if v.UID == u.ID {
				v.IsOnline = u.IsOnline
				v.UserName = u.UserName
				room.Seats[k] = v
			}
		}
	}
}

//send room base info
func (rmg *RoomMgr) SendRoomBaseInfo(u *model.User) {
	room, _ := rmg.Get(u.RoomID)
	nmsg := msg.ClientMsg{}
	nmsg.Code = MSG_ROOM_SYS_INFO_NOTICE
	nmsg.Status = SUM_MSG_STATUS
	nmsg.Msg = room.GetRoomBaseInfo()
	g_user_mgr.SendMsg(u.ID, &nmsg)
}

func (rmg *RoomMgr) BoadcastRoomsInfo(room_id int, uid uint, code int) {
	room, ok := rmg.Get(room_id)
	nmsg := msg.ClientMsg{}
	if ok {
		nmsg.Status = room.Status
		log.Debug("BoadcastRoomsInfo %d, %d,%d,%d", room_id, uid, code, room.Status)
		switch code {
		case MSG_ROOM_STATUS_CHANGE_NOTICE:
			nmsg.Code = MSG_ROOM_STATUS_CHANGE_NOTICE
			//nmsg.Status = ROOM_STATUS_USER_LOGIN
			//seat := room.GetSeatByUID(uid)
			nmsg.Msg = room.Seats
		case ROOM_STATUS_DIS_DONE:
			rmg.Clear(room_id)

			nmsg.Code = MSG_ROOM_DISSOLVED_NOTICE

		case ROOM_STATUS_OVER:
			rmg.Clear(room_id)

			nmsg.Code = MSG_ROOM_DISSOLVED_NOTICE

		case MSG_ROOM_SYS_INFO_NOTICE:
			nmsg.Code = MSG_ROOM_SYS_INFO_NOTICE
			if ROOM_STATUS_DIS_DONE == room.Status {
				rmg.Clear(room_id)
				nmsg.Msg = "解散"
			} else if ROOM_STATUS_PLAYING == room.Status {
				nmsg.Msg = room.GetRoomBaseInfo()
			} else if ROOM_STATUS_APL_DIS == room.Status {
				nmsg.Msg = room.GetRoomBaseInfo()
			} else if ROOM_STATUS_WAIT == room.Status {
				rmg.Clear(room_id)
			}
		case MSG_ROOM_SYS_DISSOLVE_NOTICE:
			nmsg.Code = MSG_ROOM_SYS_DISSOLVE_NOTICE

			nmsg.Msg = room.Seats
			if ROOM_STATUS_DIS_DONE == room.Status {
				rmg.Clear(room_id)
				nmsg.Msg = "解散"
			} else if ROOM_STATUS_PLAYING == room.Status {

			} else if ROOM_STATUS_APL_DIS == room.Status {

			}
		}
		///return room, true
		for _, v := range room.Seats {
			if v.UID != uid {
				g_user_mgr.SendMsg(v.UID, &nmsg)
			}
		}
	}
	//return nil, false
}
