package middleware

import (
	"github.com/gin-gonic/gin"
)

func direction404(c *gin.Context) {
	c.Redirect(301, "/production/list")
}
