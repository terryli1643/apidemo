package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/terryli1643/apidemo/libs/configure"
)

func init() {
	configure.Path = "../env/local/config/server.json"
}
