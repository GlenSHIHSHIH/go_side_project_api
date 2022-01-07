package backstage

import (
	errorcode "componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/pkg/errors"
)

type JwtService struct {
}

func GetJwtService() *JwtService {
	return &JwtService{}
}

//刷新 RefreshToken
func (j *JwtService) RefreshToken(refToken *backstagedto.JwtRefTokenDTO) (interface{}, error) {

	jwtInfoDTO, err := utils.ValidateAndRefreshTokenCheck(refToken.RefreshToken)
	if err != nil {
		return nil, err
	}

	//todo jwt token refresh token
	jwtToken, errT := utils.GenJwt(jwtInfoDTO.Id, jwtInfoDTO.Name)
	refreshToken, errR := utils.GenRefJwt(jwtInfoDTO.Id, jwtInfoDTO.Name)
	if errT != nil || errR != nil {
		errToken := errors.WithMessage(errors.WithStack(errT), errorcode.GENERATE_JWT_ERROR)
		errRefToken := errors.WithMessage(errors.WithStack(errR), errorcode.GENERATE_REFRESH_JWT_ERROR)
		log.Error(fmt.Sprintf("%+v,%+v", errToken, errRefToken))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.GENERATE_JWT_ERROR)
	}

	res := &backstagedto.LoginResponseDTO{
		UserInfo:     &backstagedto.JwtUserInfoDTO{Id: jwtInfoDTO.Id, Name: jwtInfoDTO.Name},
		AuthorityJwt: &backstagedto.JwtTokenDTO{Token: jwtToken, RefreshToken: refreshToken},
	}
	return res, nil
}
