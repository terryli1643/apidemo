package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/domain/model"
	"github.com/terryli1643/apidemo/libs/configure"
	"github.com/terryli1643/apidemo/libs/datasource"
	"github.com/terryli1643/apidemo/libs/logger"
	"github.com/terryli1643/apidemo/libs/orm"
	"github.com/terryli1643/apidemo/router"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func init() {
	model.InitialModels()
	configure.LoadWithJson("env/local/config/server.json")
	orm.InitCasbinEnforcer("env/local/config/casbin/rbac_model.conf")
	logger.InitLog()

}

func main() {
	datasource.GetDB()
	server := &http.Server{
		Addr:         ":9999",
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
