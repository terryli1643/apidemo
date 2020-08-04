package redirectmiddleware

import (
	"gitlab.99safe.org/Shadow/shadow-framework/logger"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("redirectmiddleware init")
}

const (
	REDIRECT_PARAMETER_KEY = "REDIRECT_PARAMETER_KEY"
)
