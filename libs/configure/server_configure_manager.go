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
	DNS          string
	Driver       string
	IdlePoolSize int
	MaxPoolSize  int
	MaxLifeTime  int64
	SqlDebug     int8
	AutoCreate   bool
}

// func NewServerConfig() {
// 	log.Infoln("ServerConfig init")
// 	LoadWithFile(file)
// }

func LoadWithJson(configFile string) {
	ServerConfig = TServerConfig{}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.WithField("file", configFile).Error("server configure init failed, file doesn't exist")
		log.Panic(err)
	}
	log.Info("Server configure", string(data))

	datajson := []byte(data)
	err = json.Unmarshal(datajson, &ServerConfig)
	if err != nil {
		log.Panic(err)
	}
	log.Infof("%+v", ServerConfig)
}
