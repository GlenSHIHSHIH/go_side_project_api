package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Direction404(c *gin.Context) {
	//使用系統預設 404
	c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
