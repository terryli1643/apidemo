package security

import (
	log "github.com/sirupsen/logrus"
)

const (
	ANONYMOUS_AUTHENTICATION_PROVIDER      = "AnonymousAuthenticationProvider"
	CASBIN_AUTHENTICATION_PROVIDER         = "CasbinAuthenticationProvider"
	DAO_AUTHENTICATION_PROVIDER            = "DaoAuthenticationProvider"
	PASSWORD_ENCODER                       = "PasswordEncoder"
	PROVIDER_MANAGER                       = "ProviderManager"
	USER_DETAILS_SERVICE                   = "UserDetailsService"
	SHADOW_SECURITY_TOKEN                  = "security_context"
	SHADOW_SECURITY_FORM_USERNAME_KEY      = "username"
	SHADOW_SECURITY_FORM_PASSWORD_KEY      = "password"
	ANONYMOUS_ROLE                         = "ROLE_ANONYMOUS"
	CASBIN_AUTHENTICATION_REQUEST_RESOLVER = "CasbinAuthenticationRequestResolver"
	USERNAME_PASSWORD_RESOLVER             = "UsernamePasswordResolver"
)

func init() {

	log.Infoln("AnonymousAuthenticationProvider init")
	RegisterAuthenticationProvider(ANONYMOUS_AUTHENTICATION_PROVIDER, newAnonymousAuthenticationProvider)

	log.Infoln("CasbinAuthenticationProvider init")
	RegisterAuthenticationProvider(CASBIN_AUTHENTICATION_PROVIDER, newCasbinAuthenticationProvider)

	log.Infoln("DaoAuthenticationProvider init")
	RegisterAuthenticationProvider(DAO_AUTHENTICATION_PROVIDER, newDaoAuthenticationProvider)

	log.Infoln("PasswordEncoder init")
	RegisterPasswordEncoder(PASSWORD_ENCODER, newDefaultPasswordEncoder)

	log.Infoln("ProviderManager init")
	RegisterAuthenticationManager(PROVIDER_MANAGER, newAuthenticationManager)

	log.Info("CasbinAuthenticationRequestResolver init")
	RegisterCasbinAuthenticationRequestResolve(CASBIN_AUTHENTICATION_REQUEST_RESOLVER, newRbacDomainCasbinAuthenticationRequestResolver)

	log.Info("UsernamePasswordResolver init")
	RegisterUsernamePasswordResolver(USERNAME_PASSWORD_RESOLVER, newFormUsernamePasswordResolver)

}
