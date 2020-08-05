package configure

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var ServerConfig TServerConfig

type TServerConfig struct {
	DataSource TDataSourceConfig
}

type TDataSourceConfig struct {
	DSN          string
	Driver       string
	IdlePoolSize int
	MaxPoolSize  int
	MaxLifeTime  int64
	SqlDebug     int8
	AutoMigrate  bool
}

func LoadWithJson(configFile string) {
	ServerConfig = TServerConfig{}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.WithField("file", configFile).Error("server configure init failed, file doesn't exist")
		log.Panic(err)
	}

	datajson := []byte(data)
	err = json.Unmarshal(datajson, &ServerConfig)
	if err != nil {
		log.Panic(err)
	}
	log.Infof("%+v", ServerConfig)
}
