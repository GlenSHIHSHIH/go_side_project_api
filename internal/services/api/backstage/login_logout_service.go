package backstage

import (
	errorCode "componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/pkg/errors"
)

type LoginLoutService struct {
}

func GetLoginLoutService() *LoginLoutService {
	return &LoginLoutService{}
}

//登入
func (u *LoginLoutService) Login(loginDTO *backstagedto.LoginDTO) (interface{}, error) {

	loginName := loginDTO.LoginName
	loginPwd := loginDTO.Password

	if len(loginName) <= 3 || len(loginPwd) <= 5 {
		otherErr := ""
		if len(loginName) <= 3 || len(loginPwd) <= 5 {
			otherErr = fmt.Sprintf(", or loginDTO.Name,loginDTO.LoginName lens <=3 and len(loginPwd) <= 5, len(loginName)=%v,len(loginPwd)=%v", len(loginName), len(loginPwd))
		}
		errData := errors.New(errorCode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v"+otherErr, errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	userService := GetUserService()
	user := userService.GetUserByLoginName(loginName)
	userCheck := false
	if user != nil {
		userCheck = userService.CheckUserPwd(loginPwd, user.Password)
	}

	if user == nil || userCheck == false {
		errData := errors.New(errorCode.UNAUTHORIZED_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.UNAUTHORIZED_ERROR_CODE, errorCode.UNAUTHORIZED_ERROR)
	}

	//todo jwt token refresh token

	return nil, nil
}
