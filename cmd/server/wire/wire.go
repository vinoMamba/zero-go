//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/vinoMamba/zero/internal/server"
	"github.com/vinoMamba/zero/pkg/log"
)

var ServerSet = wire.NewSet(server.NewHttpServer)

// var RepositorySet = wire.NewSet(
// 	repository.NewDb,
// 	repository.NewRepository,
// 	repository.NewUserRepository,
// )

// var ServiceSet = wire.NewSet(
// 	service.NewService,
// 	service.NewUserService,
// )

// var HandlerSet = wire.NewSet(
// 	handler.NewHandler,
// 	handler.NewUserHandler,
// )

func NewApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
	))
}
