package app

import (
	"bez/config"
	v1 "bez/internal/controller/http/v1"
	"bez/internal/usecase"
	"bez/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	router := chi.NewRouter()

	ga := usecase.NewGoogleAPIUseCase(cfg.CredentialsBin)

	v1.NewRouter(router, ga)

	serv := httpserver.New(router, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	_, err := os.Getgroups()

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
