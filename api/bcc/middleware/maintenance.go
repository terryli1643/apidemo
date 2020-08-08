package middleware

import (
	"github.com/gin-gonic/gin"
)

func Maintenance() gin.HandlerFunc {
	return func(c *gin.Context) {
		// paramService := service.NewParamSettingService()
		// _, enable, _ := paramService.GetParamByName("Maintenance")

		// if enable {
		// 	log.Warn("system is under maintenance")

		// 	c.AbortWithStatusJSON(http.StatusForbidden, GenericMessageBody{
		// 		Message: "系统维护",
		// 	})
		// }
	}
}
