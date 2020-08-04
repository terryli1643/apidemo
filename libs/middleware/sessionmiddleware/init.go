package sessionmiddleware

import (
	"encoding/gob"

	"gitlab.99safe.org/Shadow/shadow-framework/logger"
	"gitlab.99safe.org/Shadow/shadow-framework/security"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("session middleware init")
	gob.Register(&security.TAnonymousAuthenticationToken{})
	gob.Register(&security.TRequestAuthenticationToken{})
	gob.Register(&security.TUsernamePasswordAuthenticationToken{})
	gob.Register(&security.TWebAuthenticationDetails{})
	gob.Register(&security.TUser{})
}
