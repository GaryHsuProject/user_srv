package driver

import (
	"shop/config"

	"github.com/spf13/viper"
)

var GlobalConfig config.ServerConfig

func Init() {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.AutomaticEnv()
	if err := v.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}
}
