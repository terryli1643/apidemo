package errormiddleware

import (
	"gitlab.99safe.org/Shadow/shadow-framework/logger"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("ErrorHandler init")
}
