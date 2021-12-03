package backstage

import (
	errorcode "componentmod/internal/api/errorcode"
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
		errData := errors.New(errorcode.UNAUTHORIZED_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.UNAUTHORIZED_ERROR_CODE, errorcode.UNAUTHORIZED_ERROR)
	}

	//todo jwt token refresh token
	jwtToken, errT := utils.GenJwt(user.Id, user.Name)
	refreshToken, errR := utils.GenRefJwt(user.Id, user.Name)
	if errT != nil {
		errToken := errors.WithMessage(errors.WithStack(errT), errorcode.GENERATE_JWT_ERROR)
		errRefToken := errors.WithMessage(errors.WithStack(errR), errorcode.GENERATE_REFRESH_JWT_ERROR)
		log.Error(fmt.Sprintf("%+v,%+v", errToken, errRefToken))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.GENERATE_JWT_ERROR)
	}

	//insert login ip

	//

	res := &backstagedto.LoginResponseDTO{
		UserInfo:     &backstagedto.JwtInfoDTO{Id: user.Id, Name: user.Name},
		AuthorityJwt: &backstagedto.JwtTokenDTO{Token: jwtToken, RefreshToken: refreshToken},
	}

	return res, nil
}
