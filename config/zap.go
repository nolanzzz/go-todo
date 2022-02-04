package config

type Zap struct {
	Format   string `mapstructure:"format"`
	ShowLine bool   `mapstructure:"showline"`
}
