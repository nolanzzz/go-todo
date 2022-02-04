package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
	JWT   JWT   `mapstructure:"jwt"`
	Zap   Zap   `mapstructure:"zap"`
}
