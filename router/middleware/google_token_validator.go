package middleware

import (
	"github.com/gin-gonic/gin"
)

func GoogleTokenValidator(filterURL []string, fn func(c *gin.Context) (profileID int64, acc string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if len(c.Errors) > 0 {
		// 	return
		// }
		// paramService := service.NewParamSettingService()

		// _, enable, err := paramService.GetParamByName("GoogleTokenValidator")

		// if !enable {
		// 	Log.Warn(err)
		// 	return
		// }

		// for _, url := range filterURL {
		// 	if c.Request.URL.RequestURI() == url && c.Request.Method == http.MethodPost {
		// 		profileID, acc := fn(c)
		// 		passCode := c.GetHeader("OTP")
		// 		if passCode == "" {
		// 			c.Error(WrongGoogleToken{
		// 				Err: errors.New("google token is empty"),
		// 			})
		// 			return
		// 		}
		// 		googleTokenService := service.NewGoogleTokenService()
		// 		if ok := googleTokenService.VerifyTotp(fmt.Sprintf("%d", profileID), acc, passCode); !ok {
		// 			c.Error(WrongGoogleToken{
		// 				Err: errors.New("google token is wrong"),
		// 			})
		// 		}
		// 	}
		// }
	}
}
