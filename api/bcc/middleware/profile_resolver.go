package middleware

import (
	"github.com/gin-gonic/gin"
)

func ProfileResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
	}
}
