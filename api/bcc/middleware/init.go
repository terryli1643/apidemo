package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MENUS   = "menus"
	PROFILE = "profile"
	HANDLER = "Handler"
)

func newGenError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

func new400Error(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

func newForbiddenError(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

type GenericSuccess struct {
	//in: body
	Body GenericMessageBody
}

// swagger:model
type GenericMessageBody struct {
	Message string
}
