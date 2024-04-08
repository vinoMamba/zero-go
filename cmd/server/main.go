package main

import (
	"fmt"

	"github.com/vinoMamba/zero/pkg/config"
	"github.com/vinoMamba/zero/pkg/log"
	"go.uber.org/zap"
)

func main() {
	// init config
	conf := config.NewConfig()
	// init logger
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))

}
