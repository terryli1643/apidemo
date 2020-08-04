package loginmiddleware

import (
	"gitlab.99safe.org/Shadow/shadow-framework/logger"
	"gitlab.99safe.org/Shadow/shadow-framework/middleware"
)

var (
	Log *logger.Logger
)

const (
	LOGIN_HANDLER = "LoginHandler"
)

func init() {
	Log = logger.InitLog()
	Log.Info("LoginHandler init")
	middleware.RegisterMiddlewareHandler(LOGIN_HANDLER, newDefaultLoginHandler)
}
