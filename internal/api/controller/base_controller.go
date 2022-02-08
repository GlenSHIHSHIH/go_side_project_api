package controller

import (
	"componentmod/internal/dto"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Data interface{}

type HandlerFunc func(c *gin.Context) (Data, error)
type WebHandlerFunc func(c *gin.Context)

func WebHandler(hf WebHandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		hf(c)
	}
}

func Handler(hf HandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		data, err := hf(c)
		msg := "success"

		if err == nil {
			baseResponseDTO := &dto.BaseResponseDTO{
				Data: data,
				Msg:  msg,
			}
			c.JSON(http.StatusOK, baseResponseDTO)
		}

		if err != nil {
			//error string 切割
			code := 500
			data = nil
			msg = err.Error()
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
			c.JSON(code, baseResponseDTO)
		}
	}
}
