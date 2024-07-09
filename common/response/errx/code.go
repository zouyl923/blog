package errx

const (
	Success     = 0  //成功
	Error       = -1 //通用失败
	LoginExpire = -2 //token失效, 但refreshToken可能未失效
	LoginError  = -3 //双Token都失效
	ParamError  = -4 //参数错误

	NotFundError    = -5  //查询失败
	DeleteError     = -6  //删除失败
	UpdateError     = -7  //更新失败
	MessageError    = -8  //自定义业务错误
	UnKnowError     = -9  //未知错误
	PermissionError = -10 //权限不足
	//admin模块 101开头 01 是具体业务
	AdminNotFound = 10101 //账户或者密码错误

)
