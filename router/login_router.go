package router

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoginRouter(r *gin.Engine) {
	r.POST("/login", LoginHandler)
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Result struct {
	Msg string `json:"msg"`
}

func LoginHandler(c *gin.Context) {
	dump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug(string(dump))

	data := Login{}
	err = c.Bind(&data)
	if err != nil {
		newClientError(c, err)
		return
	}
	log.Debugf("data is :: %+v", data)
	c.JSON(200, Result{Msg: "Success"})
}
