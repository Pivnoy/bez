package app

import (
	"bez/config"
	v1 "bez/internal/controller/http/v1"
	"bez/internal/usecase"
	"bez/internal/usecase/repo"
	"bez/pkg/httpserver"
	"bez/pkg/postgres"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	usRp := repo.NewUserRepo(pg)
	flRp := repo.NewFileTorrentRepo(pg)
	srRp := repo.NewServiceRepo(pg)

	ga := usecase.NewGoogleAPIUseCase(cfg.CredentialsBin)
	dr := usecase.NewDriveAPI()
	us := usecase.NewUserUseCase(usRp)
	fl := usecase.NewFileUseCase(flRp)
	sr := usecase.NewServiceUseCase(srRp)

	ld := usecase.NewLoadUseCase(fl, sr, ga, dr)
	err = ld.Preload(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	v1.NewRouter(router, ga, dr, us, sr, fl)

	serv := httpserver.New(router, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

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
