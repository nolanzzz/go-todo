package config

type Config struct {
	JWT    JWT    `mapstructure:"jwt"`
	Zap    Zap    `mapstructure:"zap"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	System System `mapstructure:"system"`
}
