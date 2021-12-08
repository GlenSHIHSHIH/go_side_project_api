package middleware

import (
	"github.com/gin-gonic/gin"
)

func authorityJwtMenuCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header().
		// utils.ValidateAndTokenCheck()

		c.Next()

		// // after request
		// latency := time.Since(t)
		// log.Print(latency)

		// // access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)
	}
}
