package service

import (
	"time"
	"todo/global"
)

type JwtService struct{}

var JwtServiceApp JwtService

func (j *JwtService) SetRedisJWT(username string, jwt string) (err error) {
	expire := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(username, jwt, expire).Err()
	return err
}

func (j *JwtService) GetRedisJWT(username string) (jwt string, err error) {
	jwt, err = global.REDIS.Get(username).Result()
	return jwt, err
}
