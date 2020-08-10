package configure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	LoadWithJson("../../env/local/config/server.json")
}

func TestServerConfig(t *testing.T) {
	serverConfigure := ServerConfig

	assert.NotEmpty(t, serverConfigure.DataSource.IdlePoolSize)
	assert.NotEmpty(t, serverConfigure.DataSource.MaxLifeTime)
	assert.NotEmpty(t, serverConfigure.DataSource.MaxPoolSize)
	assert.NotEmpty(t, serverConfigure.DataSource.DSN)
}
