package middleware

import (
	"componentmod/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func jwtValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authorization")
		token := strings.Replace(bearerToken, "Bearer ", "", 1)

		//驗證使用者 jwt token
		jwtInfoDTO, err := utils.ValidateAndTokenCheck(token)

		//登入的頁面權限

		if err == nil {
			c.Set("userInfo", jwtInfoDTO)
			c.Next()
		} else {
			middlewareErrHandler(c, nil, err)
		}
	}
}
