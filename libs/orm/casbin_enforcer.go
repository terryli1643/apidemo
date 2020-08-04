package orm

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/datasource"
)

var e *casbin.Enforcer

func InitCasbinEnforcer(config string) {
	log.Info("Casbin Enforcer init")
	a, _ := gormadapter.NewAdapterByDB(datasource.GetDB())
	e, _ = casbin.NewEnforcer(config, a)
	e.LoadPolicy()
}

func GetCasbinEnforcer() *casbin.Enforcer {
	if e == nil {
		panic("Casbin has not initial")
	}
	return e
}
