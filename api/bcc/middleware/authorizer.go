package middleware

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/terryli1643/apidemo/libs/logger"
	"github.com/terryli1643/apidemo/service"
)

// Authorizer is a gin midlleware for authorizer request operation
func Authorizer(context string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}

		v, ok := c.Get("jwt_claims")
		if !ok {
			err := service.NewCasbinAuthService().Authenticate(service.CasbinPolicy{
				Sub:     "ROLE_ANONYMOUS",
				Domain:  context,
				Obj:     c.Request.URL.Path,
				Act:     c.Request.Method,
				Service: "",
				Eft:     "",
			})
			if err != nil {
				log.Error(err)
				newGenError(c, err.Error())
				return
			}
		} else {
			claims := v.(jwt.MapClaims)
			userID := claims["id"].(int64)
			// 验证session是否过期
			sessionService := service.NewSessionService()
			if sessionService.CheckSessionExpired(userID) {
				log.Warning("Session已过期")
				err := service.NewCasbinAuthService().Authenticate(service.CasbinPolicy{
					Sub:     "ROLE_ANONYMOUS",
					Domain:  context,
					Obj:     c.Request.URL.Path,
					Act:     c.Request.Method,
					Service: "",
					Eft:     "",
				})
				if err != nil {
					log.Error(err)
					newGenError(c, "Session已过期, "+err.Error())
					return
				}
				return
			} else {
				err := sessionService.SetSessionTime(userID)
				if err != nil {
					log.Error(err)
					newGenError(c, err.Error())
					return
				}
			}
			//验证权限
			err := service.NewCasbinAuthService().Authenticate(service.CasbinPolicy{
				Sub:     fmt.Sprint(userID),
				Domain:  context,
				Obj:     c.Request.URL.Path,
				Act:     c.Request.Method,
				Service: "",
				Eft:     "",
			})
			if err != nil {
				log.Error(err)
				newGenError(c, err.Error())
				return
			}
		}

	}
}
