//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/liqieli/raiden/internal/server"
	"github.com/liqieli/raiden/pkg/app"
	"github.com/liqieli/raiden/pkg/log"
	"github.com/spf13/viper"
)

var taskSet = wire.NewSet(server.NewTask)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		taskSet,
		newApp,
	))
}
