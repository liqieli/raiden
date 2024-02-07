//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/liqieli/raiden/internal/repository"
	"github.com/liqieli/raiden/internal/server"
	"github.com/liqieli/raiden/pkg/app"
	"github.com/liqieli/raiden/pkg/log"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

// build App
func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		server.NewMigrate,
		newApp,
	))
}
