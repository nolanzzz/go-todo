package config

type Database struct {
	Seed    bool `mapstructure:"seed"`
	Migrate bool `mapstructure:"migrate"`
}
