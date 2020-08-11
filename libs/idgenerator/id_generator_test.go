package idgenerator

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/terryli1643/apidemo/libs/configure"
)

func init() {
	log.Infoln("id_generator_test init")
	configure.Path = "../../env/local/config/server.json"
}

func TestGenerateStringID(t *testing.T) {
	gen := NewIdGenerator()
	assert.NotNil(t, gen)
	for i := 0; i < 2000; i++ {
		id := gen.GenerateStringID("user")
		t.Log(id)
		assert.NotEmpty(t, id)
	}
}

func TestGenerateRandID(t *testing.T) {
	gen := NewIdGenerator()
	assert.NotNil(t, gen)
	t.Log(gen.GenerateGuardID("user"))
}

func TestMessID(t *testing.T) {
	key := "0ZV6AWDFGSCU9HJL578X1MBKN24QERTYIOP3"

	gen := NewIdGenerator()
	assert.NotNil(t, gen)
	l := gen.GenerateLongID("user")
	t.Logf("long : %d", l)

	s := gen.ChaosID(l, key)
	t.Logf("string : %s", s)

	sl := gen.RestoreID(s, key)
	t.Logf("long : %d", sl)
}
