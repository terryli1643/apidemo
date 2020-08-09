package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
			token := c.MustGet("jwt_token").(string)
			// 验证session是否过期
			sessionService := service.NewSessionService()
			err := sessionService.SetSessionExpireTime(token, service.DefaultMaxLifeTime)
			if err != nil {
				log.Warning("Session已过期", err)
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
			}
			//验证权限
			claims := v.(jwt.MapClaims)
			err = service.NewCasbinAuthService().Authenticate(service.CasbinPolicy{
				Sub:     claims["Account"].(string),
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
