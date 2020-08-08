package middleware

// func JWTNoSessionParse(secret string, maxLifeTime int64) gin.HandlerFunc {
// remitList := map[string]string{
// 	"/bcc/login": http.MethodPost,
// 	"/mcc/login": http.MethodPost,
// 	//	"/mcc/translateRecords": http.MethodPost,
// 	"/papi/login":           http.MethodPost,
// 	"/qmapi/login":          http.MethodPost,
// 	"/papi/smsCode":         http.MethodPost,
// 	"/papi/register":        http.MethodPost,
// 	"/papi/indexHeroesList": http.MethodGet,
// 	"/papi/forgetPassword":  http.MethodPost,
// }
// return func(c *gin.Context) {
// 	Log.Debugf("jwt processing %s %s", c.Request.URL.Path, c.Request.Method)
// 	if len(c.Errors) > 0 {
// 		return
// 	}

// 	if v, ok := remitList[c.Request.URL.Path]; ok && c.Request.Method == v {
// 		anonymousToken := security.NewAnonymousAuthenticationToken()
// 		anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
// 			RemoteAddress: c.ClientIP(),
// 			RequestURI:    c.Request.URL.Path,
// 		})
// 		c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
// 		return
// 	}

// 	// 获取token
// 	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
// 		b := ([]byte(secret))
// 		return b, nil
// 	})

// 	if err != nil {
// 		Log.Warn("jwt token error", err)
// 		anonymousToken := security.NewAnonymousAuthenticationToken()
// 		anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
// 			RemoteAddress: c.ClientIP(),
// 			RequestURI:    c.Request.URL.Path,
// 		})
// 		c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
// 		return
// 	}

// 	// 校验并解析token
// 	var account string
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		Log.Debug(claims)
// 		account = claims["Account"].(string)
// 		// 创建权限认证token
// 		authRequest := security.NewUsernamePasswordAuthenticationToken(account, "")
// 		authRequest.SetAuthenticated(true)
// 		c.Set(security.SHADOW_SECURITY_TOKEN, authRequest)
// 		c.Set("jwt_token", token.Raw)
// 		c.Set("jwt_claims", claims)

// 		if err != nil {
// 			Log.Error(err)
// 			c.AbortWithError(http.StatusUnauthorized, err)
// 		}

// 	} else {
// 		anonymousToken := security.NewAnonymousAuthenticationToken()
// 		anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
// 			RemoteAddress: c.ClientIP(),
// 			RequestURI:    c.Request.URL.Path,
// 		})
// 		c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
// 		return
// 	}

// }
// }
