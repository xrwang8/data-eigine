package cmd

import (
	"context"
	conf "data-engine/config"
	"data-engine/protocol"
	"data-engine/routers"
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
	"os/signal"
	"syscall"
)

type StartCommand struct{}

var ServeCmd = &cobra.Command{
	Use:          "serve",
	Short:        "data-engine",
	SilenceUsage: true,
	Long:         "data-engine Start service",
	Run: func(c *cobra.Command, args []string) {
		logger := conf.GetLogger()
		opts := fx.Options(
			fx.WithLogger(func() fxevent.Logger {
				return logger.GetFxLogger()
			}),
			fx.Invoke(NewStartCommand().Run()),
		)
		ctx := context.Background()
		app := fx.New(CommonModules, opts)
		err := app.Start(ctx)
		defer app.Stop(ctx)
		if err != nil {
			logger.Fatal(err)
		}
	},
}

func (s *StartCommand) Run() conf.CommandRunner {
	return func(conf conf.Config, router conf.RequestHandler, route routers.Routes, logger conf.Logger, database conf.Database) {
		route.Setup()
		logger.Info("Running server")
		serviceManager := newServiceManager(conf, router)
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go serviceManager.WaitStop(ch)
		if err := serviceManager.Start(); err != nil {
			logger.Fatalf("Service startup failed %s", err.Error())
		}
	}
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

type seviceManager struct {
	http *protocol.HttpService
}

func (s *seviceManager) Start() error {
	return s.http.Start()
}

// Process interrupt signals from outside, such as Terminal
func (s *seviceManager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			fmt.Printf("received signal: %s", v)
			s.http.Stop()
		}
	}
}

func newServiceManager(conf conf.Config, router conf.RequestHandler) *seviceManager {
	return &seviceManager{
		http: protocol.NewHttpService(conf, router),
	}
}
