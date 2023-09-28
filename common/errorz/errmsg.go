package errorz

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	//用户模块
	message[EMAIL_SEND_ERROR] = "邮件发送失败,请稍后再试"
	message[CAPTCHA_VALIDATE_ERROR] = "验证码错误"
	message[PASSWORD_VALIDATE_ERROR] = "密码错误"
	//todo模块
	message[SEND_WARNING_ERROR] = "添加待办事项提醒失败"
	message[PAST_DEADLINE_ERROR] = "无效日期"
	message[OPERATE_CRON_ERROR] = "添加提醒失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
