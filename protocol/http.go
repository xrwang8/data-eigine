package protocol

import (
	"context"
	conf "data-engine/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HttpService struct {
	server *http.Server
}

func NewHttpService(conf conf.Config, router conf.RequestHandler) *HttpService {
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              fmt.Sprintf(":%s", conf.App.Port),
		Handler:           router.Gin,
	}
	return &HttpService{
		server: server,
	}
}

func (s *HttpService) Start() error {
	log.Printf("[info] start http server listening %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil

}

func (s *HttpService) Stop() {
	fmt.Println("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("shut down http service error, %s", err)
	}
}

func (s *HttpService) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			fmt.Printf("received signal: %s", v)
			s.Stop()
		}
	}
}
