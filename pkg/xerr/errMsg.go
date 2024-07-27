package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COxyzON_ERROR] = "SERVER_ERROR"
	message[REUQEST_PARAM_ERROR] = "REUQEST_PARAM_ERROR"
	message[TOKEN_EXPIRE_ERROR] = "TOKEN_EXPIRE_ERROR"
	message[TOKEN_GENERATE_ERROR] = "TOKEN_GENERATE_ERROR"
	message[DB_ERROR] = "DB_ERROR"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "DB_UPDATE_AFFECTED_ZERO_ERROR"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "SERVER_ERROR"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
