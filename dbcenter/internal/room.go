package internal

import (
	. "qpgame/base"
	//. "qpgame/base/function"
	"qpgame/base/model"
	"qpgame/msg"
	"time"

	"github.com/name5566/leaf/log"
)

type Room struct {
	ID            int
	Name          string
	PeopleNumbers int //how man people
	CreatedAt     time.Time
	OverAt        time.Time
	GameNumbers   int
	CreatedBy     uint
	Seats         map[int]*Seat
	Status        int //0-wait 1-playing 2-apply dissolve 3-disssolved 4-over
	//create_by *User
	Game          IGame
	ZhuangIndex   int
	AplDisUserID  uint
	PlaiedNumbers int
}

type Seat struct {
	IDX             int
	UID             uint
	UserName        string
	IsOnline        int
	IsReday         int //0-no 1-yes
	IsAgreeDissolve int //0-init 1-yes  2-no

	Score int
}

func (r *Room) CreateRoomID() int {
	return 10010
}
func (r *Room) Join(u *model.User) (*Room, bool) {
	seat := new(Seat)

	//
	idx := -1
	for i := 0; i < r.PeopleNumbers; i++ {
		if _, ok := r.Seats[i]; ok {
			continue
		}

		if idx == -1 {
			idx = i
		} else {
			break;
		}
		
	}
	
	seat.IDX = idx
	seat.UID = u.ID
	seat.UserName = u.UserName
	seat.IsOnline = u.IsOnline
	r.Seats[seat.IDX] = seat

	u.RoomID = r.ID

	return r, true
}

func (r *Room) SetReady(u *model.User) {
	var ready_seat_num int = 0
	for _, v := range r.Seats {
		if v.UID == u.ID {
			v.IsReday = 1
			ready_seat_num++
			continue
		}

		if v.IsReday == 1 {
			ready_seat_num++
			continue
		}
	}

	if r.PeopleNumbers == ready_seat_num {
		r.Status = ROOM_STATUS_PLAYING
		r.Begin()
	}

}


func (r *Room) Begin() {
	//game_id := r.CreateGameID()
	log.Debug("room Begin........ %d.......", r.PlaiedNumbers)
	seats := make(map[int]uint, len(r.Seats))
	for k, v := range r.Seats {
		seats[k] = v.UID
	}

	r.Game.Init(1, seats)
	r.Game.Begin()
	r.Game.Shuffle()
	r.Game.Deal()
	
	r.SysGameInfo2Client(false)
}

func (r *Room) Leave(u *model.User) {
	for k, v := range r.Seats {
		if v.UID == u.ID {
			delete(r.Seats, k)
			u.RoomID = 0
		}
	}
}

func (r *Room) Dissolve(u *model.User) {
	log.Debug("room Dissolve........")
	r.Status = ROOM_STATUS_APL_DIS
	r.AplDisUserID = u.ID
	for _, v := range r.Seats {
		if v.UID == u.ID {
			v.IsAgreeDissolve = 1
		}
	}
}

func (r *Room) DoDissolve() {
	log.Debug("room DoDissolve........")
	r.Status = ROOM_STATUS_DIS_DONE
	r.Clear()
	//notice to client
}

func (r *Room) UnDissolve() {
	log.Debug("room UnDissolve........")
	r.Status = ROOM_STATUS_DIS_DONE

	//notice to client
}
func (r *Room) Clear() {
	for _, v := range r.Seats {
		u := g_user_mgr.Get(v.UID)
		if u != nil {
			u.RoomID = 0
		}
	}
}
func (r *Room) AgreeDissolve(u *model.User) bool {
	var agree_count int = 0
	var un_agree_count = 0
	for _, v := range r.Seats {
		if v.IsAgreeDissolve == 1 {
			agree_count++
		} else {
			if v.UID == u.ID {
				v.IsAgreeDissolve = 1
				agree_count++
			} else {
				un_agree_count++
			}
		}
	}

	if agree_count+un_agree_count == r.PeopleNumbers {
		if agree_count >= un_agree_count {
			r.DoDissolve()
		} else {
			r.UnDissolve()
		}
	}
	return false
}

func (r *Room) GetSeatByUID(uid uint) *Seat {
	for _, s := range r.Seats {
		if s.UID == uid {
			return s
		}
	}

	return nil
}

func (r *Room) GetRoomBaseInfo() *msg.RoomBaseInfo {
	rbf := new (msg.RoomBaseInfo)
	rbf.ID = r.ID
	rbf.Name = r.Name
	rbf.PeopleNumbers = r.PeopleNumbers
	rbf.GameNumbers = r.GameNumbers
	rbf.CreatedBy = r.CreatedBy
	rbf.Seats = make(map[int]*msg.SeatBaseInfo, 0)
	rbf.Status = r.Status
	rbf.ZhuangIndex = r.ZhuangIndex
	rbf.AplDisUserID = r.AplDisUserID
	rbf.PlaiedNumbers = r.PlaiedNumbers
	
	for k,v := range r.Seats {
		se := new(msg.SeatBaseInfo)
		se.IDX = v.IDX
		se.UID = v.UID
		se.UserName = v.UserName
		se.IsOnline = v.IsOnline
		se.IsReday = v.IsReday
		se.IsAgreeDissolve = v.IsAgreeDissolve
		se.Score = v.Score
		rbf.Seats[k] = se
	}

	return rbf
}

//游戏信息
func (r *Room) GetGameBaseInfo(is_show bool, uid uint) *msg.GameInfo{
	game_info := new(msg.GameInfo)
	game_info.Init()

	ret := r.Game.GetDeskBaseInfo(is_show).(map[string][]int)
	
	game_info.Cards = ret["cards"]
	

	seats_row := r.Game.GestSeatsCards(is_show, uid).([]map[string][]int)
	for k, v := range seats_row {
		seat_cards := new(msg.GameSeatCard)
		seat_cards.IDX = k
		seat_cards.Init()
		seat_cards.Holds = v["holds"]
		seat_cards.Folds = v["folds"]
		seat_cards.AnGangs = v["an_gangs"]
		seat_cards.DianGangs = v["dian_gangs"]
		seat_cards.CaGangs = v["ca_gangs"]
		seat_cards.Pengs = v["pengs"]

		game_info.Seats[k] = seat_cards
	}

	return game_info	
}

func (r *Room) SysGameInfo2Client(is_show bool) {
	nmsg := msg.ClientMsg{}
	nmsg.Code  = MSG_GAME_BEGIN
	
	///return room, true
	for k, v := range r.Seats {
		game_info  := r.GetGameBaseInfo(is_show, v.UID)
		game_info.SelfIndex = k
		nmsg.Msg = game_info
		g_user_mgr.SendMsg(v.UID, &nmsg)
	}
}
func (r *Room) SendGameBaseInfo(is_show bool, uid uint) {	
}