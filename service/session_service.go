package service

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/terryli1643/apidemo/libs/redisclient"
)

const (
	SessionKey = "session:online"
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
