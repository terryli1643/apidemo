package service

import (
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/redisclient"
)

const DefaultMaxLifeTime = 300

type SessionService struct {
}

var sessionServiceObj *SessionService

func NewSessionService() *SessionService {
	if sessionServiceObj == nil {
		l.Lock()
		if sessionServiceObj == nil {
			sessionServiceObj = &SessionService{}
		}
		l.Unlock()
	}
	return sessionServiceObj
}

//CreateJWT 创建登陆令牌
func (service SessionService) CreateJWT(clainms map[string]interface{}, hmacSampleSecret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(clainms))
	tokenString, err = token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		log.Error(err)
	}
	return
}

//Createsession 创建session
func (service SessionService) CreateSession(token string, userID int64, ip string) error {
	conn := redisclient.GetConn()
	defer conn.Close()
	key := fmt.Sprintf("session:%s:%d:", token, userID)

	log.WithFields(log.Fields{
		"key":    key,
		"userID": userID,
		"IP":     ip,
	}).Debug("Create Session")

	err := conn.Send("HSET", key, "userID", userID, "IP", ip)
	if err != nil {
		log.Error(err)
		return err
	}
	err = conn.Send("EXPIRE", key, 3600)
	if err != nil {
		log.Error(err)
		return err
	}
	err = conn.Flush()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

//DeleteSessionByToken 根据令牌删除token
func (service SessionService) DeleteSessionByToken(token string) error {
	conn := redisclient.GetConn()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%s:*", token)))
	if err != nil {
		log.Error(err)
		return err
	}
	if reply != nil && len(reply) > 0 {
		key := reply[0]
		conn.Do("DEL", key)
	}
	return nil
}

//DeleteSessionByUserID 根据用户ID删除session
func (service SessionService) DeleteSessionByUserID(userID int64) error {
	conn := redisclient.GetConn()
	defer conn.Close()

	//删除session
	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:*:%d:*", userID)))
	if err != nil {
		log.Error(err)
		return err
	}
	if reply != nil && len(reply) > 0 {
		key := reply[0]
		conn.Do("DEL", key)
	}
	return nil
}

//GetSessionByUserID 根据用户ID获取session 的token
func (service SessionService) GetTokenByUserID(userID int64) (token string, err error) {
	conn := redisclient.GetConn()
	defer conn.Close()

	//删除session
	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:*:%d:*", userID)))
	if err != nil {
		log.Error(err)
		return "", err
	}
	if reply != nil && len(reply) > 0 {
		token = reply[0]
		token = strings.Split(token, ":")[1]
	}
	return
}

//SetSessionExpireTime 延期session过期时间
func (service SessionService) SetSessionExpireTime(token string, maxLifeTime int64) error {
	conn := redisclient.GetConn()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%s:*:", token)))
	if err != nil {
		log.Error(err)
		return err
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session过期时间
		conn.Do("EXPIRE", key, maxLifeTime)
		return nil
	}
	err = errors.New("token expired")
	log.Error(err)
	return err
}
