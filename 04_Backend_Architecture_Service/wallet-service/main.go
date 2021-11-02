package main

import (
	"log"
	"wallet-service/config"
	"wallet-service/data/datasource/postgres"
	"wallet-service/data/repository"
	"wallet-service/domain/usecase"
	"wallet-service/server"
	"wallet-service/service"

	"github.com/kelseyhightower/envconfig"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func main() {
	var cfg config.Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error loading environment config: %v", err)
	}

	db, err := postgres.NewWalletDB(&cfg)
	if err != nil {
		log.Fatalf("Error loading environment config: %v", err)
	}

	r := repository.NewWalletRepository(db)
	r.Initialize()
	u := usecase.NewWalletUsecase(r)
	s := service.NewWalletService(u)

	server, _ := server.NewGrpcServer(cfg.ServerPort, s)
	server.Run()
}
