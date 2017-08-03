package base

const (
	//client
	MSG_CLIENT = 0

	MSG_CLIENT_LOGIN_REQ = MSG_CLIENT + 1 //user login
	MSG_CLIENT_LOGIN_RSP = MSG_CLIENT + 2

	MSG_CLIENT_USERINFO_REQ = MSG_CLIENT + 3 //user info
	MSG_CLIENT_USERINFO_RSP = MSG_CLIENT + 4

	MSG_CLIENT_RECONNECT_REQ = MSG_CLIENT + 5 //reconnect
	MSG_CLIENT_RECONNECT_RSP = MSG_CLIENT + 6

	MSG_CLIENT_LOGOUT_REQ = MSG_CLIENT + 7 //user logout
	MSG_CLIENT_LOGOUT_RSP = MSG_CLIENT + 8

	MSG_CLIENT_REPEAT_NOTICE = MSG_CLIENT + 9 //repeat landing

	MSG_CLIENT_TOKEN_EXPIRED_NOTICE = MSG_CLIENT + 10 //token expired

	/////////////////////////////////dbcenter/////////////////////////
	MSG_DBCENTER = 2000

	MSG_DBCENTER_LOGIN_NOTICE       = MSG_DBCENTER + 1 //user login sys notice
	MSG_DBCENTER_LOGOUT_NOTICE      = MSG_DBCENTER + 2 //user logout sys notice
	MSG_DB_CENTER_DISCONNECT_NOTICE = MSG_DBCENTER + 3 //user disconnect sys notice

	///////////////////////////////room msg////////////////////////////////////////////////
	MSG_ROOM            = 3000
	MSG_ROOM_CREATE_REQ = MSG_ROOM + 1 //create romm request
	MSG_ROOM_CREATE_RSP = MSG_ROOM + 2 //

	MSG_ROOM_JOIN_REQ = MSG_ROOM + 3
	MSG_ROOM_JOIN_RSP = MSG_ROOM + 4

	MSG_ROOM_APPLY_DISSOLVE_REQ = MSG_ROOM + 5
	MSG_ROOM_APPLY_DISSOLVE_RSP = MSG_ROOM + 6

	MSG_ROOM_AGREE_DISSOLVE_REQ = MSG_ROOM + 7
	MSG_ROOM_AGREE_DISSOLVE_RSP = MSG_ROOM + 8

	MSG_ROOM_DISSOLVED_NOTICE = MSG_ROOM + 9  //notice room is dissolve
	MSG_ROOM_SYS_INFO_NOTICE  = MSG_ROOM + 10 //room info

	MSG_ROOM_LEAVE_REQ = MSG_ROOM + 11 //user leave
	MSG_ROOM_LEAVE_RSP = MSG_ROOM + 12 //

	MSG_ROOM_STATUS_CHANGE_NOTICE = MSG_ROOM + 13
	MSG_ROOM_SYS_DISSOLVE_NOTICE = MSG_ROOM + 14

	MSG_ROOM_REJOIN_SYS_NOTICE = MSG_ROOM + 15

	//////////////////////Game Msg///////////////////////////////
	MSG_GAME_BASE  = 4000
	MSG_GAME_BEGIN = MSG_GAME_BASE + 1

	//////////////////////////////////////////////////////////////////////////////
	SUM_MSG_STATUS         = 0
	SUM_MSG_STATUS_SUCCESS = SUM_MSG_STATUS
	SUM_MSG_STATUS_ERROR   = SUM_MSG_STATUS + 1
	SUM_MSG_STATUS_REPEAT  = SUM_MSG_STATUS + 2
	SUM_MSG_STATUS_TO_MUCH = SUM_MSG_STATUS + 3

	/////////////////////////////sub room status
	//0-wait 1-playing 2-apply dissolve 3-disssolved 4-over 5-room not exist 6-user login 7 user logout
	ROOM_STATUS_BASE      = 0
	ROOM_STATUS_WAIT      = ROOM_STATUS_BASE
	ROOM_STATUS_PLAYING   = ROOM_STATUS_BASE + 1
	ROOM_STATUS_APL_DIS   = ROOM_STATUS_BASE + 2
	ROOM_STATUS_DIS_DONE  = ROOM_STATUS_BASE + 3
	ROOM_STATUS_OVER      = ROOM_STATUS_BASE + 4
	ROOM_STATUS_NOT_EXIST = ROOM_STATUS_BASE + 5
	//ROOM_STATUS_USER_LOGIN  = ROOM_STATUS_BASE + 6
	//ROOM_STATUS_USER_LOGOUT = ROOM_STATUS_BASE + 7
)
