package config

import (
	"flag"
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yaml", "config path")
		flag.Parse()
	}

	if envConf == "" {
		envConf = "config/local.yaml"
	}

	return getConfig(envConf)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
