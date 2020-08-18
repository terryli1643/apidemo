package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	log "github.com/terryli1643/apidemo/libs/logger"
	"github.com/terryli1643/apidemo/libs/redisclient"
)

const (
	SessionKey       = "session:online"
	HmacSampleSecret = "wErUOtNOXiPHVPunb9Y0tn$KmatydruRTKlaUdup9newmb9Y0du$2a$10"
)

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

type IUserDetailService interface {
	LoadUserByUsername(username string) (IUserDetails, error)
}

type IUserDetails interface {
	GetID() int64
	GetUsername() string
	GetPassword() string
	IsAccountLocked() bool
}

func (service *SessionService) Login(username string, password string, userService IUserDetailService) (token string, userDetail IUserDetails, err error) {
	userDetail, err = userService.LoadUserByUsername(username)
	if err != nil {
		err = errors.New("用户名密码错误")
		log.Error(err)
		return
	}
	passwordService := NewPasswordSerice()
	ok := passwordService.Matches(password, userDetail.GetPassword())
	if !ok {
		err = errors.New("用户名密码错误")
		log.Error(err)
		return
	}

	sessionService := NewSessionService()
	token, err = sessionService.CreateSession(userDetail.GetID(), userDetail.GetUsername())
	if err != nil {
		log.Error(err)
		return
	}
	return
}

//CreateJWT 创建登陆令牌
func (service SessionService) CreateJWT(claims map[string]interface{}) (tokenString string, err error) {
	claims["timestamp"] = time.Now().Nanosecond()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	tokenString, err = token.SignedString([]byte(HmacSampleSecret))
	if err != nil {
		log.Error(err)
	}
	return
}

func (service SessionService) CreateSession(userID int64, account string) (token string, err error) {
	claim := map[string]interface{}{
		"id":      userID,
		"account": account,
	}
	token, err = service.CreateJWT(claim)
	if err != nil {
		log.Error(err)
		return
	}
	err = service.SetSessionTime(userID)
	if err != nil {
		log.Error(err)
	}
	return
}

//根据用户ID删除session
func (service SessionService) DeleteSession(userID int64) error {
	conn := redisclient.GetConn()
	defer conn.Close()

	_, err := conn.Do("ZREM", SessionKey, userID)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (service SessionService) CheckSessionExpired(userID int64) (ok bool) {
	conn := redisclient.GetConn()
	defer conn.Close()

	_, err := redis.Int64(conn.Do("ZRANK", SessionKey, userID))
	if err != nil {
		ok = false
	} else {
		ok = true
	}
	return
}

//记录付客账号在线的心跳时间
func (service SessionService) SetSessionTime(userID int64) (err error) {
	//redis 设置最后在线时间, 使用有序集合存储payerID和时间戳
	conn := redisclient.GetConn()
	defer conn.Close()

	//生成时间戳
	timestamp := time.Now().Unix()
	log.Debug(timestamp)
	_, err = conn.Do("zadd", SessionKey, timestamp, fmt.Sprint(userID))
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//移除心跳过期的账号，并将其设置为掉线状态
func (service SessionService) ResetSessionTime(sessionTime int64) (err error) {
	conn := redisclient.GetConn()
	defer conn.Close()

	reply, err := redis.Int64s(conn.Do("ZRANGEBYSCORE", SessionKey, 0, sessionTime))

	if err != nil {
		log.Error(err)
		return err
	}

	for _, v := range reply {
		if err != nil {
			log.Error(err)
			return err
		}

		_, err := conn.Do("ZREM", SessionKey, v)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
