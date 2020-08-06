package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/terryli1643/apidemo/libs/security"
)

func ProfileResolver(callback func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
		token := c.MustGet(security.SHADOW_SECURITY_TOKEN)
		if secureToken, ok := token.(*security.TUsernamePasswordAuthenticationToken); ok && secureToken.IsAuthenticated() {
			callback(c)
		}
	}
}
