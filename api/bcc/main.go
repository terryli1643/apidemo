package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/api/bcc/router"
	"github.com/terryli1643/apidemo/domain/model"
	"github.com/terryli1643/apidemo/libs/configure"
	"github.com/terryli1643/apidemo/libs/orm"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func init() {
	configure.Path = "env/local/config/server.json"
	model.InitialModels()
	orm.InitCasbinEnforcer("env/local/config/casbin/rbac_model.conf")
}

func main() {
	server := &http.Server{
		Addr:         configure.New().BccServer.Port,
		Handler:      router.MainRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
