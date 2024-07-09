package errx

var cn = map[int]string{
	Success:         "操作成功！",
	Error:           "操作失败！",
	LoginExpire:     "登录过期",
	LoginError:      "登录失效",
	ParamError:      "参数错误",
	NotFundError:    "没有此信息",
	DeleteError:     "删除失败",
	UpdateError:     "更新失败",
	MessageError:    "自定义错误",
	UnKnowError:     "未知错误",
	PermissionError: "权限不足",
	AdminNotFound:   "账号或者密码错误！",
}

func GetCnMessage(code int) string {
	msg, ok := cn[code]
	if !ok {
		msg = cn[Error]
	}
	return msg
}
