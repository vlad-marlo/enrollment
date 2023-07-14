package main

import (
	"github.com/vlad-marlo/enrollment/internal/config"
	"github.com/vlad-marlo/enrollment/internal/controller"
	"github.com/vlad-marlo/enrollment/internal/controller/http"
	"github.com/vlad-marlo/enrollment/internal/pkg/logger"
	pg "github.com/vlad-marlo/enrollment/internal/pkg/pgx"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx/client"
	"github.com/vlad-marlo/enrollment/internal/service"
	"github.com/vlad-marlo/enrollment/internal/store/pgx"
	"go.uber.org/fx"
)

const serversGroup = `group:"servers"`

//	@title		Enrollment service
//	@version	1.0

func main() {
	fx.New(NewApp()).Run()
}

// NewApp prepares fx options to configure and start application.
func NewApp() fx.Option {
	return fx.Options(
		fx.Provide(
			logger.New,
			AsServer(http.New),
			fx.Annotate(config.NewControllerConfig, fx.As(new(controller.Config))),
			fx.Annotate(service.New, fx.As(new(controller.Service))),
			fx.Annotate(client.New, fx.As(new(pg.Client))),
			fx.Annotate(pgx.New, fx.As(new(service.Repository))),
		),
		fx.Invoke(
			fx.Annotate(RunServers, fx.ParamTags(serversGroup)),
		),
	)
}

// AsServer annotates the given constructor to state that
// it provides a route to the servers group.
func AsServer(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(controller.Interface)),
		fx.ResultTags(serversGroup),
	)
}

func RunServers(servers []controller.Interface, lc fx.Lifecycle) {
	for _, srv := range servers {
		lc.Append(fx.Hook{
			OnStart: srv.Start,
			OnStop:  srv.Stop,
		})
	}
}
