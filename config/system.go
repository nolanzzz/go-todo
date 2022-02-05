package config

import "fmt"

type System struct {
	Port        int  `mapstructure:"port"`
	UseRedisJWT bool `mapstructure:"use-redis-jwt"`
}

func (s *System) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}
