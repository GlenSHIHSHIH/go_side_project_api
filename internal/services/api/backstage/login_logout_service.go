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

	userService := GetUserService()
	user := userService.GetUserByLoginName(loginDTO.LoginName)
	userCheck := false
	if user != nil {
		userCheck = userService.CheckUserPwd(loginDTO.Password, user.Password)
	}

	if user == nil || userCheck == false {
		errData := errors.New(errorCode.UNAUTHORIZED_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.UNAUTHORIZED_ERROR_CODE, errorCode.UNAUTHORIZED_ERROR)
	}

	//todo jwt token refresh token

	return nil, nil
}
