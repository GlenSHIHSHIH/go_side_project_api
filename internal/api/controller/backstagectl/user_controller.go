package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/validator.v2"
)

var (
	UserCreate = controller.Handler(CreateUser)
	// UserLogin  = controller.Handler(Login)
)

// @Summary Backstage UserLogin
// @Success 200 {json} json
// @Router /backstage/user/create [post]
func CreateUser(c *gin.Context) (controller.Data, error) {
	var user *model.User
	err := c.Bind(&user)
	err = validator.Validate(user)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	userService := backstage.GetUserService()
	return userService.CreateUser(user)
}
