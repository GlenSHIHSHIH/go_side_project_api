package errorcode

const (
	PARAMETER_ERROR_CODE    = 400
	UNAUTHORIZED_ERROR_CODE = 401
	SERVER_ERROR_CODE       = 500
)
const (
	PARAMETER_ERROR                = "參數錯誤"
	PARAMETER_LOGIN_NAME_DUPLICATE = "登入帳號重複"
	UNAUTHORIZED_ERROR             = "身份驗證錯誤"
	REFRESH_AUTHORIZED_ERROR       = "刷新身份錯誤"
	SERVER_ERROR                   = "程式錯誤"
	GENERATE_JWT_ERROR             = "產生 jwt 錯誤"
	GENERATE_REFRESH_JWT_ERROR     = "產生 Refresh jwt 錯誤"
)
