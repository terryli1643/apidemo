package service

import (
	"errors"

	"github.com/terryli1643/apidemo/libs/orm"
)

type CasbinAuthService struct {
}

var casbinAuthServiceObj *CasbinAuthService

func NewCasbinAuthService() *CasbinAuthService {
	if casbinAuthServiceObj == nil {
		l.Lock()
		if casbinAuthServiceObj == nil {
			casbinAuthServiceObj = &CasbinAuthService{}
		}
		l.Unlock()
	}
	return casbinAuthServiceObj
}

type CasbinPolicy struct {
	Sub     string
	Domain  string
	Obj     string
	Act     string
	Service string
	Eft     string
}

func (provider CasbinAuthService) Authenticate(policy CasbinPolicy) (err error) {
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

	if ok, err := orm.GetCasbinEnforcer().Enforce(param...); !ok {
		if err != nil {
			log.Error(err)
		}
		err = errors.New("权限校验失败")
	}
	return err
}
