package internal

//. "qpgame/base"
//. "qpgame/base/function"
//"qpgame/base/model"
//"qpgame/msg"
//"qpgame/base/model"
//"math/rand"
//"time"

//"github.com/name5566/leaf/log"

type IGame interface {
	Begin()
	Init(int, map[int]uint)
	Shuffle()
	MoPai(bool) (int, bool)
	Deal()
	GetDeskBaseInfo(bool) interface{}
	GestSeatsCards( bool, uint) interface{} 
}
