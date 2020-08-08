package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/terryli1643/apidemo/libs/configure"
)

func init() {
	log.Infoln("service_test init")
	configure.LoadWithJson("../env/local/config/server.json")
}
