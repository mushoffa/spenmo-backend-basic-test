package main

import (
	"log"
	"user-service/config"
	"user-service/data/datasource/postgres"
	"user-service/data/repository"
	"user-service/domain/usecase"
	"user-service/service"

	"github.com/kelseyhightower/envconfig"
	"github.com/mushoffa/go-library/grpc"
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
		log.Fatalf("Error loading environment config: %v", err)
	}

	r := repository.NewUserRepository(db)
	u := usecase.NewUserUsecase(r)
	s := service.NewUserService(u)

	server := grpc.NewGrpcServer(cfg.ServerPort)
}
