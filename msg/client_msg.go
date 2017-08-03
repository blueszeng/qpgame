package msg

type ClientMsg struct {
	Code   int
	Status int //0-success 1-failed
	Msg    interface{}
}

type UserInfo struct {
	ID        uint
	UserName  string
	IsOnline  int
	RoomID    int
	RoomCards int
	Status    int
	Nickname  string
	Type      int
}

type RoomBaseInfo struct {
	ID            int
	Name          string
	PeopleNumbers int //how man people
	GameNumbers   int
	CreatedBy     uint
	Seats         map[int]*SeatBaseInfo
	Status        int //0-wait 1-playing 2-apply dissolve 3-disssolved 4-over
	ZhuangIndex   int
	AplDisUserID  uint
	PlaiedNumbers int //paly counts
}

type SeatBaseInfo struct {
	IDX             int
	UID             uint
	UserName        string
	IsOnline        int
	IsReday         int //0-no 1-yes
	IsAgreeDissolve int //0-init 1-yes  2-no
	Score int
}

type GameInfo struct {
	Cards []int
	SelfIndex int
	Seats map[int]*GameSeatCard
}

type GameSeatCard struct {
	IDX int

	Holds     []int  
	Folds     []int

	AnGangs   []int
	DianGangs []int
	CaGangs   []int
	Pengs     []int
}

func (gi *GameInfo) Init() {
	gi.Seats = make(map[int]*GameSeatCard, 0)
}

func (gs *GameSeatCard) Init() {
	gs.Holds = make([]int, 0)  
	gs.Folds = make([]int, 0)

	gs.AnGangs = make([]int, 0)
	gs.DianGangs = make([]int, 0)
	gs.CaGangs = make([]int, 0)
	gs.Pengs = make([]int, 0)
}