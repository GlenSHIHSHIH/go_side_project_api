package backstage

import (
	errorcode "componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/log"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type LoginLogoutService struct {
}

func GetLoginLogoutService() *LoginLogoutService {
	return &LoginLogoutService{}
}

//登入
func (l *LoginLogoutService) Login(loginDTO *backstagedto.LoginDTO, ip string) (interface{}, error) {

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
	if errT != nil || errR != nil {
		errToken := errors.WithMessage(errors.WithStack(errT), errorcode.GENERATE_JWT_ERROR)
		errRefToken := errors.WithMessage(errors.WithStack(errR), errorcode.GENERATE_REFRESH_JWT_ERROR)
		log.Error(fmt.Sprintf("%+v,%+v", errToken, errRefToken))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.GENERATE_JWT_ERROR)
	}

	//set login ip and time
	// user.LoginIP = utils.GetLocalIP()
	user.LoginIP = ip
	user.LoginTime.Time = time.Now()
	sqldb := db.GetMySqlDB()
	sqldb.Save(&user)

	res := &backstagedto.LoginResponseDTO{
		UserInfo:     &backstagedto.JwtUserInfoDTO{Id: user.Id, Name: user.Name},
		AuthorityJwt: &backstagedto.JwtTokenDTO{Token: jwtToken, RefreshToken: refreshToken},
	}

	return res, nil
}

//登出
func (l *LoginLogoutService) Logout(userInfo *backstagedto.JwtUserInfoDTO) (interface{}, error) {

	//移除  user navigation (menus) cache
	menuService := GetMenuService()
	err := menuService.RemoveMenuCache(userInfo.Id)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.SERVER_ERROR)
		log.Error(fmt.Sprintf("logout error: %+v", errData))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SERVER_ERROR)
	}
	//移除 其他cache ...

	return nil, nil
}
