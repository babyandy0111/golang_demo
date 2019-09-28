package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "請求參數錯誤",
	ERROR_EXIST_TAG:                "已存在data",
	ERROR_NOT_EXIST_TAG:            "該data不存在",
	ERROR_NOT_EXIST_ARTICLE:        "該data不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token失效",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超時",
	ERROR_AUTH_TOKEN:               "Token生成失敗",
	ERROR_AUTH:                     "Token錯誤",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
