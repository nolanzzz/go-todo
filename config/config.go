package config

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql"`
	JWT    JWT    `mapstructure:"jwt"`
	System System `mapstructure:"system"`
	Zap    Zap    `mapstructure:"zap"`
}
