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
	Username string `form:"Username" json:"Username" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
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
		c.Error(err)
		return
	}
	log.Debugf("data is :: %+v", data)
	c.JSON(200, Result{Msg: "Success"})
}
