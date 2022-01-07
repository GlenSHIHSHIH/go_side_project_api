package validate

import (
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils"

	"github.com/gin-gonic/gin"
)

func UserInfoValidate(c *gin.Context) (*backstagedto.JwtUserInfoDTO, error) {

	JwtInfoDTO, exist := c.Get("userInfo")

	if exist == false || JwtInfoDTO == nil {
		err := utils.CreateApiErr(errorcode.UNAUTHORIZED_ERROR_CODE, errorcode.UNAUTHORIZED_ERROR)
		return nil, err
	}

	return JwtInfoDTO.(*backstagedto.JwtUserInfoDTO), nil
}
