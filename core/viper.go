package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"todo/global"
)

func Viper() *viper.Viper {
	config := "config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	v.WatchConfig() // watch config changes
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			panic(fmt.Errorf("fatal error loading configs: %s \n", err))
		}
		fmt.Println("config file changed:", e.Name)
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("fatal error loading configs: %s \n", err))
	}
	return v
}
