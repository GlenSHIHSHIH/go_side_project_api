package backstage

import (
	"componentmod/internal/api/controller"
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	UserCreate = controller.Handler(CreateUser)
	// UserLogin  = controller.Handler(Login)
)

// @Summary Backstage check password
// @Success 200 {json} json
// @param pwd path string true "pwd"
// @Router /backstage/login [post]
// func Login(c *gin.Context) (controller.Data, error) {
// 	userService := backstage.GetUserService()
// 	return datalist, nil
// }

// @Summary Backstage UserLogin
// @Success 200 {json} json
// @Router /backstage/user/create [post]
func CreateUser(c *gin.Context) (controller.Data, error) {
	var user *model.User
	err := c.Bind(&user)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorCode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	userService := backstage.GetUserService()
	return userService.CreateUser(user)
}
