package backstage

import (
	errorCode "componentmod/internal/api/errorcode"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db/model"
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
func (u *LoginLoutService) Login(user *model.User) (interface{}, error) {

	if err != nil || len(user.Name) <= 3 || len(user.LoginName) <= 3 || len(userPwd) <= 5 {

		errData := errors.WithMessage(errors.WithStack(err), errorCode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v"+otherErr, errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	userService := GetUserService()
	userService.GetUserByLoginName()
	userService.CheckUserPwd()
	return nil, nil
}
