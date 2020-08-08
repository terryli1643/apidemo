package middleware

import (
	"github.com/gin-gonic/gin"
)

// UsernamePasswordLoginFilter is a gin midlleware for resolve user login
func UsernamePasswordLoginFilter(loginPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if c.Request.URL.RequestURI() == loginPath {
		// 	if c.Request.Method == http.MethodPost {
		// 		usernamePasswordResolver := security.UsernamePasswordResolverInstance(security.USERNAME_PASSWORD_RESOLVER)
		// 		username, password := usernamePasswordResolver.ObtainUsernamePassword(c)
		// 		authRequest := security.NewUsernamePasswordAuthenticationToken(username, password)
		// 		authRequest.SetDetails(&security.TWebAuthenticationDetails{
		// 			RemoteAddress: c.ClientIP(),
		// 			RequestURI:    c.Request.URL.Path,
		// 		})

		// 		authentication := security.AuthenticationManagerInstance(security.PROVIDER_MANAGER).Authenticate(authRequest)

		// 		Log.WithFields(logrus.Fields{
		// 			"Authenticated": authentication.IsAuthenticated(),
		// 			"Details":       authentication.GetDetails(),
		// 			"Principal":     authentication.GetPrincipal(),
		// 		}).Debug("UsernamePasswordLoginFilter")

		// 		if authentication.IsAuthenticated() == false {
		// 			details := authentication.GetDetails()
		// 			if details == nil {
		// 				err := security.WrongUserNamePasswordError{
		// 					Err: errors.New("wrong username or wrong password"),
		// 				}
		// 				c.Error(err)
		// 				return
		// 			}
		// 			if userDetails, ok := details.(security.IUserDetails); ok {
		// 				var err error
		// 				if userDetails.IsAccountExpired() {
		// 					err = security.AccountExpiredError{
		// 						Err: errors.New("account is expired"),
		// 					}
		// 				} else if userDetails.IsAccountLocked() {
		// 					err = security.AccountLockedError{
		// 						Err: errors.New("account is locked"),
		// 					}
		// 				} else if userDetails.IsCredentialsExpired() {
		// 					err = security.WrongUserNamePasswordError{
		// 						Err: errors.New("password is expired"),
		// 					}
		// 				} else {
		// 					err = security.WrongUserNamePasswordError{
		// 						Err: errors.New("wrong username or wrong password"),
		// 					}
		// 				}
		// 				c.Error(err)
		// 				return
		// 			}
		// 		} else {
		// 			c.Set(security.SHADOW_SECURITY_TOKEN, authentication)
		// 		}
		// 	}
		// }

		// c.Next()
	}
}
