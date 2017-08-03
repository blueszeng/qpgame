package conf

import (
	"log"
	"time"
)

var (
	// log conf
	//LogFlag = log.LstdFlags
	LogFlag = log.Lshortfile
	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000

	//db
	DB_HOST = "127.0.0.1"
	DB_PORT = 3306
	DB_USER = "root"
	DB_PWD  = "sqlmima"
	DB_NAME = "lf_game"
)
