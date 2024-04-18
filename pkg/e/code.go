package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	// 动画相关错误
	ERROR_NOT_EXIST_ANIMATION = 10001

	// 认证相关错误
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	// 数据库错误
	ERROR_DB = 30001
)