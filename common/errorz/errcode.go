package errorz

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	SERVER_COMMON_ERROR uint32 = 10001 + iota
	REUQEST_PARAM_ERROR
	DB_ERROR
)

// 用户模块
const (
	CAPTCHA_VALIDATE_ERROR uint32 = 20001 + iota
	EMAIL_SEND_ERROR
	PASSWORD_VALIDATE_ERROR
)

// todo模块
const (
	SEND_WARNING_ERROR uint32 = 30001 + iota
	PAST_DEADLINE_ERROR
	OPERATE_CRON_ERROR
)
