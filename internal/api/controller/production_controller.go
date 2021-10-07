package controller

import (
	"componentmod/internal/utils"

	"github.com/gin-gonic/gin"
)

var (
	ProductionList = Handler(Production)
)

func Production(c *gin.Context) (Data, error) {

	return nil, utils.CreateApiErr(118881, "未定義錯誤")
}
