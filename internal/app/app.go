package app

import (
	"bez/config"
	v1 "bez/internal/controller/http/v1"
	"bez/internal/usecase"
	"bez/internal/usecase/repo"
	"bez/pkg/httpserver"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	pg := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)

	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	usRp := repo.NewUserRepo(db)

	ga := usecase.NewGoogleAPIUseCase(cfg.CredentialsBin)
	dr := usecase.NewDriveAPI()
	us := usecase.NewUserUseCase(usRp)

	router := gin.New()
	v1.NewRouter(router, ga, dr, us)

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
