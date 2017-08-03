package internal

//. "qpgame/base"
//. "qpgame/base/function"
//"qpgame/base/model"
//"qpgame/msg"
	
import (
	"github.com/name5566/leaf/log"
	"math/rand"
	"time"
)

type GameKWX struct {
	ID        int
	Cards     [84]int
	LastIndex int
	GameIndex int
	GameSeats map[int]*GameSeat
	Turn      int
}

type GameSeat struct {
	UID       uint
	Holds     []int
	Folds     []int
	AnGangs   []int
	DianGangs []int
	CaGangs   []int
	Pengs     []int

	CountMap map[int]int
	TingMap  map[int]int

	CanGang   int
	CanPeng   int
	CanHu     int
	CanChuPai int
	CanGangai int

	IsHud  int
	HuInfo interface{}
}

func (g *GameKWX) Init(game_id int, seat map[int]uint) {
	g.Turn = 0
	g.ID = game_id
	g.Cards = [84]int{}
	g.LastIndex = 0
	g.GameIndex = len(g.Cards) - 1
	g.GameSeats = make(map[int]*GameSeat)
	for k, v := range seat {
		s := new(GameSeat)
		s.Holds = make([]int, 0)
		s.Folds = make([]int, 0)
		s.AnGangs = make([]int, 0)
		s.DianGangs = make([]int, 0)
		s.CaGangs = make([]int, 0)
		s.Pengs = make([]int, 0)
		s.CountMap = make(map[int]int, 0)
		s.TingMap = make(map[int]int, 0)

		s.UID = v
		g.GameSeats[k] = s
	}
}

func (g *GameKWX) Begin() {
	/*for _, s := range g.GameSeats {
		s.Holds = make([]int, 0)
		s.Folds = make([]int, 0)
		s.AnGangs = make([]int, 0)
		s.DianGangs = make([]int, 0)
		s.CaGangs = make([]int, 0)
		s.Pengs = make([]int, 0)
		s.CountMap = make(map[int]int, 0)
		s.TingMap = make(map[int]int, 0)
	}*/

	mj := 0
	//tong
	for i := 1; i < 10; i++ {
		for j := 0; j < 4; j++ {
			g.Cards[mj] = i
			mj++
		}
	}
	//tiao
	for i := 11; i < 20; i++ {
		for j := 0; j < 4; j++ {
			g.Cards[mj] = i
			mj++
		}
	}
	for i := 21; i < 24; i++ {
		for j := 0; j < 4; j++ {
			g.Cards[mj] = i
			mj++
		}
	}

	log.Debug("majiang=%v", g.Cards)
	return
}

func (g *GameKWX) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbs := len(g.Cards)
	for i := 0; i <numbs; i++ {
		temp := g.Cards[i]
		g.Cards[i] = g.Cards[r.Intn(numbs)]
		g.Cards[r.Intn(numbs)] = temp
	}

	log.Debug("Shuffle majiang=%v", g.Cards)
}

func (g *GameKWX) Deal() {
	for k, seat := range g.GameSeats {
		for i := 0; i < 13; i++ {
			pai, _ := g.MoPai(true)
			//seat.Holds[i] = pai
			seat.Holds = append(seat.Holds, pai)
		}

		if g.Turn == k {
			pai, _ := g.MoPai(true)
			//seat.Holds[i] = pai
			seat.Holds = append(seat.Holds, pai)
		}
	}
	
}


func (g *GameKWX) MoPai(direction bool) (int, bool) {
	if (g.LastIndex > g.GameIndex) {
		return 0, false
	}
	if (direction) {
		pai := g.Cards[g.LastIndex]
		g.LastIndex++
		return pai, true
	} else {
		pai := g.Cards[g.GameIndex]
		g.GameIndex--
		return pai, true
	}
}

func (g *GameKWX) GetDeskBaseInfo(is_show bool)  interface{} {
	ret := make(map[string][]int, 0)
	numbs := g.GameIndex - g.LastIndex + 1
	cards := make([]int, numbs)
	j := 0
	for i := g.GameIndex; i <= g.LastIndex; i++ {
		if is_show {
			cards[j] = g.Cards[i]
		} else {
			cards[j] = 0
		}
	}
	ret["cards"] = cards
	return ret
}

func (g *GameKWX) GestSeatsCards(is_show bool, uid uint) interface{} {
	ret := make([]map[string][]int, 3)
	for k, v := range g.GameSeats {
		seat := make(map[string][]int, 0)
		holds := make([]int, len(v.Holds))
		folds := make([]int, len(v.Folds))
		an_gangs := make([]int, len(v.AnGangs))
		dian_gangs := make([]int, len(v.DianGangs))
		ca_gangs := make([]int, len(v.CaGangs))
		pengs := make([]int, len(v.Pengs))
		
		if (is_show) {
			holds = v.Holds
			folds = v.Folds
			an_gangs = v.AnGangs
			dian_gangs = v.DianGangs
			ca_gangs = v.CaGangs
			pengs = v.Pengs
			
		} else {
			if uid == v.UID {
				holds = v.Holds
				an_gangs = v.AnGangs
			}
			
			folds = v.Folds
			dian_gangs = v.DianGangs
			ca_gangs = v.CaGangs
			pengs = v.Pengs
		}
		seat["holds"] = holds
		seat["folds"] = folds
		seat["an_gangs"] = an_gangs
		seat["dian_gangs"] = dian_gangs
		seat["ca_gangs"] = ca_gangs
		seat["pengs"] = pengs
		ret[k] = seat
	}

	return ret
}