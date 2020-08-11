package configure

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Path = "env/local/config/server.json"
var ServerConfig *TServerConfig

type TServerConfig struct {
	DataSource  TDataSourceConfig
	RedisServer TRedisConfiugre
	LogRotate   TRotateFileConfig
	BccServer   TBccServer
}

type TDataSourceConfig struct {
	DSN          string
	Driver       string
	IdlePoolSize int
	MaxPoolSize  int
	MaxLifeTime  int64
	AutoMigrate  bool
}

type TRedisConfiugre struct {
	Password    string
	Host        string
	Port        string
	DB          string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type TRotateFileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
}

type TBccServer struct {
	Port    string
	Context string
}

func New() *TServerConfig {
	if ServerConfig == nil {
		loadWithJson()
	}
	return ServerConfig
}

func loadWithJson() {
	data, err := ioutil.ReadFile(Path)
	if err != nil {
		log.Printf("server configure init failed, file doesn't exist, %s", Path)
		log.Panic(err)
	}

	datajson := []byte(data)
	err = json.Unmarshal(datajson, &ServerConfig)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("%+v", ServerConfig)
}
