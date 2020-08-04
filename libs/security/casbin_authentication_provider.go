package security

import (
	"github.com/casbin/casbin/v2"
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/orm"
)

var (
	casbinAuthenticationProvider *TCasbinAuthenticationProvider
)

// TCasbinAuthenticationProvider an implementation that retrieves user details from a UserDetailService
type TCasbinAuthenticationProvider struct {
	enforcer *casbin.Enforcer
}

type TCasbinPolicyDetails struct {
	Sub     string
	Domain  string
	Obj     string
	Act     string
	Service string
	Eft     string
}

func newCasbinAuthenticationProvider() IAuthenticationProvider {
	if casbinAuthenticationProvider == nil {
		return &TCasbinAuthenticationProvider{
			enforcer: orm.GetCasbinEnforcer(),
		}
	}
	return casbinAuthenticationProvider
}

func (provider *TCasbinAuthenticationProvider) Authenticate(authentication IAuthentication) IAuthentication {
	if requestAuthenticationToken, ok := authentication.(*TRequestAuthenticationToken); ok {
		details := authentication.GetDetails()
		if policy, ok := details.(TCasbinPolicyDetails); ok {
			var param []interface{}

			if policy.Sub != "" {
				param = append(param, policy.Sub)
			}
			if policy.Domain != "" {
				param = append(param, policy.Domain)
			}
			if policy.Obj != "" {
				param = append(param, policy.Obj)
			}
			if policy.Act != "" {
				param = append(param, policy.Act)
			}
			if policy.Service != "" {
				param = append(param, policy.Service)
			}
			if policy.Eft != "" {
				param = append(param, policy.Eft)
			}
			if ok, _ := provider.enforcer.Enforce(param...); ok {
				requestAuthenticationToken.SetAuthenticated(true)
			} else {
				log.WithField("param:", param).Debug("without promission")
			}
		}
		return authentication
	}
	return nil
}
