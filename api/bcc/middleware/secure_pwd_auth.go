package middleware

// func SecurePwdAuth(secret string) gin.HandlerFunc {
// secureMethod := map[string]string{
// 	"/acc/secure/token":               http.MethodGet,
// 	"/acc/player":                     http.MethodPatch,
// 	"/acc/agent/loginpassword":        http.MethodPatch,
// 	"/acc/agent/securepassword":       http.MethodPatch,
// 	"/acc/globalagent/loginpassword":  http.MethodPatch,
// 	"/acc/globalagent/securepassword": http.MethodPatch,
// 	"/acc/player/loginPassword":       http.MethodPatch,
// 	"/acc/player/securePassword":      http.MethodPatch,
// 	"/acc/financialAccount":           http.MethodGet,
// 	"/pcc/secure/token":               http.MethodGet,
// 	"/gcc/secure/token":               http.MethodGet,
// }
// return func(c *gin.Context) {
// 	if len(c.Errors) > 0 {
// 		return
// 	}

// 	if v, ok := secureMethod[c.Request.URL.Path]; ok && v == c.Request.Method {
// 		secureToken := c.GetHeader("Authorization2")

// 		if secureToken == "" {
// 			c.AbortWithStatus(http.StatusPreconditionFailed)
// 			return
// 		}
// 		token, err := jwt.Parse(secureToken, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(secret), nil
// 		})
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusPreconditionFailed)
// 			return
// 		}
// 		if !token.Valid {
// 			c.AbortWithStatus(http.StatusPreconditionFailed)
// 			return
// 		}

// 	}
// }
// }
