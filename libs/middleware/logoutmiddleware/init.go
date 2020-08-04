package logoutmiddleware

import (
	"github.com/astaxie/beego/session"
	"gitlab.99safe.org/Shadow/shadow-framework/logger"
	"gitlab.99safe.org/Shadow/shadow-framework/middleware"
)

var (
	Log            *logger.Logger
	globalSessions *session.Manager
)

const (
	LOGOUT         = "logout"
	LOGOUT_HANDLER = "LogoutHandler"
)

func init() {
	Log = logger.InitLog()
	Log.Info("DefaultLogoutUrlRegistry init")
	middleware.RegisterMiddlewareHandler(LOGOUT_HANDLER, newDefaultLogoutHandler)
}
