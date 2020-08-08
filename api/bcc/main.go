package main

import (
	"net/http"
	"time"

	"github.com/terryli1643/apidemo/api/bcc/router"
	"github.com/terryli1643/apidemo/domain/model"
	"github.com/terryli1643/apidemo/libs/configure"
	"github.com/terryli1643/apidemo/libs/logger"
	"github.com/terryli1643/apidemo/libs/orm"
	"golang.org/x/sync/errgroup"
)

var (
	g   errgroup.Group
	log = logger.New()
)

func init() {
	model.InitialModels()
	configure.LoadWithJson("env/local/config/server.json")
	orm.InitCasbinEnforcer("env/local/config/casbin/rbac_model.conf")

}

func main() {
	server := &http.Server{
		Addr:         configure.ServerConfig.BccServer.Port,
		Handler:      router.MainRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
