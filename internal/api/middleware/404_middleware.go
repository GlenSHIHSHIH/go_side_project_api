package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func direction404(c *gin.Context) {
	c.JSON(http.StatusNotFound, "")
}
