package config

type Database struct {
	Migrate bool `mapstructure:"migrate"`
}
