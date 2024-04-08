// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/vinoMamba/zero/internal/server"
	"github.com/vinoMamba/zero/pkg/log"
)

// Injectors from wire.go:

func NewApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	engine := server.NewHttpServer(logger)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewHttpServer)