package security

import (
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/orm"
)

func init() {
	log.Infoln("security init")
	config.LoadWithJson("../../env/local/config/server.json")
	orm.InitCasbinEnforcer("../../env/local/config/casbin/rbac_model.conf")
}