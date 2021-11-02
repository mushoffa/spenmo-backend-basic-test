package main

import (
	"card-service/config"
	"card-service/data/datasource/postgres"
	"card-service/data/repository"
	"card-service/domain/usecase"
	"card-service/server"
	"card-service/service"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func main() {
	var cfg config.Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error loading environment config: %v", err)
	}

	db, err := postgres.NewCardDB(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := repository.NewCardRepository(db)
	r.Initialize()
	u := usecase.NewCardUsecase(r)
	s := service.NewCardService(u)

	server, _ := server.NewGrpcServer(cfg.ServerPort, s)

	server.Run()
}
