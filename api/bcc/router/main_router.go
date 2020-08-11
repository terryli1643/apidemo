package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/terryli1643/apidemo/api/bcc/middleware"
	"github.com/terryli1643/apidemo/libs/configure"
)

func MainRouter() http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ConcurrentLimit(20, 5*time.Minute))
	r.Use(middleware.JWT())
	r.Use(middleware.Authorizer(configure.New().BccServer.Context))
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	g := r.Group(configure.New().BccServer.Context)
	LoginRouter(g)
	return r
}
