package middleware

import (
	"github.com/gin-gonic/gin"
)

// VisitSettingFilter is a gin midlleware for Visit request operation
func VisitSettingFilter(site string, filterFn func(site string, ip string) (bool, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if len(c.Errors) > 0 {
		// 	return
		// }
		// ip := sutils.GetRealIp(c.Request)
		// ok, err := filterFn(site, ip)
		// if err != nil {
		// 	Log.Error(err)
		// 	c.AbortWithError(http.StatusForbidden, err)
		// 	return
		// }
		// if !ok {
		// 	Log.Warnf("visit ip forbidden : %v", ip)
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }

		// // c.Next()
	}
}
