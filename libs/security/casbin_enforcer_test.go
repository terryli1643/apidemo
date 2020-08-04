package security

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/terryli1643/apidemo/libs/orm"
)

func TestCasbinFunc(t *testing.T) {
	enforcer := orm.GetCasbinEnforcer()
	roles := enforcer.GetAllRoles()
	assert.NotEmpty(t, roles)
	for _, role := range roles {
		log.Debug(role)
	}
}

func TestCasbinCheck(t *testing.T) {
	enforcer := orm.GetCasbinEnforcer()
	result, _ := enforcer.Enforce("test", "/", "/bcc/login", "POST")
	role, _ := enforcer.GetRolesForUser("test")
	log.Debug(role)
	assert.True(t, result)
}
