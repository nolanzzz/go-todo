package config

type Redis struct {
	DB             int    `mapstructure:"db"`
	Addr           string `mapstructure:"addr"`
	Password       string `mapstructure:"password"`
	KeyRankTodos   string `mapstructure:"key-rank-todos"`
	KeyRankMinutes string `mapstructure:"key-rank-minutes"`
}
