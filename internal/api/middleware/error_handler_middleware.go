package middleware

import (
	"componentmod/internal/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrHandlerMiddleware(c *gin.Context, data interface{}, err error) {

	//error string 切割
	code := 500
	msg := err.Error()
	errData := strings.Split(err.Error(), "@")
	if len(errData) >= 2 {
		errCode, _ := strconv.Atoi(errData[0])
		code = errCode
		msg = errData[1]
	}
	baseResponseDTO := &dto.BaseResponseDTO{
		Data: data,
		Msg:  msg,
	}
	c.AbortWithStatusJSON(code, baseResponseDTO)
}
