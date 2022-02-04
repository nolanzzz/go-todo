package config

import "fmt"

type System struct {
	Port int `mapstructure:"port"`
}

func (s *System) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}
