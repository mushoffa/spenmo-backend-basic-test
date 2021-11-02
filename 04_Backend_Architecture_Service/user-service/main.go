package main

import (
	"log"
	"user-service/config"
	"user-service/data/datasource/postgres"
	"user-service/data/repository"
	"user-service/domain/usecase"
	"user-service/server"
	"user-service/service"

	"github.com/kelseyhightower/envconfig"
	// "github.com/mushoffa/go-library/server/grpc"
	// "github.com/mushoffa/spenmo-proto/protos"
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

	db, err := postgres.NewUserDB(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := repository.NewUserRepository(db)
	r.Initialize()
	u := usecase.NewUserUsecase(r)
	s := service.NewUserService(u)

	server, _ := server.NewGrpcServer(cfg.ServerPort, s)

	server.Run()
}
