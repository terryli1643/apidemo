package router

import (
	"github.com/gin-gonic/gin"
)

func LoginRouter(g *gin.RouterGroup) {
	g.POST("/login", LoginHandler)
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Result struct {
	Msg string `json:"msg"`
}

func LoginHandler(c *gin.Context) {
	log.Debug("11111111111111111")
	data := Login{}
	err := c.Bind(&data)
	if err != nil {
		newClientError(c, err)
		return
	}
	c.JSON(200, Result{Msg: "Success"})
}
