package service

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/terryli1643/apidemo/domain/model"
)

func TestCreateJWT(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserType": fmt.Sprint(model.UserTypeAdmin),
		"UserID":   "1",
	})
	tokenString, err := token.SignedString([]byte("WErUOtNOXiPRTKlaUdup9newmb9HVPunb9YuY0du$2a$100tn$Kmatydr"))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(tokenString)
}
