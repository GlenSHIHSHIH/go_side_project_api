package forestage

import (
	"componentmod/internal/api/controller"
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/services/api/forestage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ProductionRank = controller.Handler(GetProductionRank)
)

func GetProductionRank(c *gin.Context) (controller.Data, error) {
	count, err := strconv.ParseInt(c.Param("count"), 10, 0)

	if err != nil {
		log.Error(fmt.Sprintf("err:%+v", err))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	rankService := forestage.GetRankService()
	return rankService.GetProductionRank(int(count))
}
