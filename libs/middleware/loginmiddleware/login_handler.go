package loginmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.99safe.org/Shadow/shadow-framework/middleware"
)

//TDefaultLoginHandler login middlewareHandler implementation
type TDefaultLoginHandler struct{}

func newDefaultLoginHandler() middleware.IMiddlewareHandler {
	return new(TDefaultLoginHandler)
}

//Handle handle redirect
func (handler *TDefaultLoginHandler) Handle(c *gin.Context) {
	Log.Debugln("loginHandler redirect to /")
	c.Redirect(http.StatusFound, "/")
}
