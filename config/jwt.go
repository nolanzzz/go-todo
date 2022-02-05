package config

type JWT struct {
	ExpiresTime int64  `mapstructure:"expires-time"`
	SigningKey  string `mapstructure:"signing-key"`
}
