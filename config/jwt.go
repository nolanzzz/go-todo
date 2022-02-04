package config

type JWT struct {
	ExpiresTime int64  `mapstructure:"expires-time"`
	Key         string `mapstructure:"key"`
}
