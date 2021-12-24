package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Direction404(c *gin.Context) {
	c.JSON(http.StatusNotFound, "")
}
