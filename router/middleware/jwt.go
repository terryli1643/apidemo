package middleware

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func JWT(secret string, maxLifeTime int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("jwt processing %s %s", c.Request.URL.Path, c.Request.Method)
		if len(c.Errors) > 0 {
			return
		}

		// 获取token
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			log.Warn("jwt token error", err)
			return
		}

		// 校验并解析token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Debug(claims)
			if err != nil {
				log.Error(err)
				c.AbortWithError(http.StatusUnauthorized, err)
			}

			c.Set("jwt_token", token.Raw)
			c.Set("jwt_claims", claims)
		}

	}
}
