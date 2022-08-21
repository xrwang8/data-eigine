package cmd

import (
	"context"
	conf "data-engine/config"
	"data-engine/routers"
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StartCommand struct {
	server *http.Server
}

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
		s.server = s.ConstructorServer(conf, router)
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go s.WaitStop(ch)
		s.Start()
	}
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (s *StartCommand) ConstructorServer(conf conf.Config, router conf.RequestHandler) *http.Server {
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              fmt.Sprintf(":%s", conf.App.Port),
		Handler:           router.Gin,
	}
	return server
}

func (s *StartCommand) Start() {
	log.Printf("[info] start http server listening %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Fatal(http.ErrServerClosed)
		}
		log.Fatalf("start service error, %s", err.Error())
	}

}

func (s *StartCommand) Stop() {
	fmt.Println("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("shut down http service error, %s", err)
	}
}

func (s *StartCommand) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			fmt.Printf("received signal: %s", v)
			s.Stop()
		}
	}
}
