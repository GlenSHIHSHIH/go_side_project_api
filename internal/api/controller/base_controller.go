package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Data interface{}

type HandlerFunc func(c *gin.Context) (Data, error)

func Handler(hf HandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		respMap := gin.H{}
		data, err := hf(c)
		code := 200
		msg := "success"

		if err != nil {
			//error string 切割
			code = 500
			data = nil
			msg = err.Error()
			errData := strings.Split(err.Error(), "@")
			if len(errData) >= 2 {
				errCode, _ := strconv.Atoi(errData[0])
				code = errCode
				msg = errData[1]
			}
		}

		respMap["code"] = code
		respMap["data"] = data
		respMap["msg"] = msg
		c.JSON(http.StatusOK, respMap)
	}
}
