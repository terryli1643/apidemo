package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terryli1643/apidemo/service"
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
	data := Login{}
	err := c.Bind(&data)
	if err != nil {
		err = errors.New("用户名密码错误")
		newClientError(c, err)
		return
	}

	adminService := service.NewAdminService()
	admin, err := adminService.GetAdminByLoginName(data.Username)
	if err != nil {
		err = errors.New("用户名密码错误")
		newClientError(c, err)
		return
	}

	sessionService := service.NewSessionService()
	token, err := sessionService.CreateSession(admin.ID, admin.Account)
	if err != nil {
		newClientError(c, err)
		return
	}
	c.JSON(http.StatusOK, CommonResp{
		Code:    http.StatusOK,
		Message: "success",
		Body:    token,
	})
}
