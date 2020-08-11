package service

import (
	log "github.com/terryli1643/apidemo/libs/logger"
	"golang.org/x/crypto/bcrypt"
)

var passwordService *PasswordService

type PasswordService struct{}

func NewPasswordSerice() *PasswordService {
	if adminServiceObj == nil {
		l.Lock()
		if adminServiceObj == nil {
			passwordService = new(PasswordService)
		}
		l.Unlock()
	}
	return passwordService
}

func (passwordEncoder *PasswordService) Encode(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(hash)
}

func (passwordEncoder *PasswordService) Matches(password string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
	return err == nil
}
