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
	client "github.com/mushoffa/go-library/server/grpc"
	"github.com/mushoffa/spenmo-proto/protos"
	"google.golang.org/grpc"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
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

	userClient, err := client.NewGrpcClient(cfg.UserClientURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to User GRPC Client: %s", err)
	}

	walletClient, err := client.NewGrpcClient(cfg.WalletClientURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Wallet GRPC Client: %s", err)
	}

	userClientService := protos.NewUserServiceClient(userClient.Conn)
	walletClientService := protos.NewWalletServiceClient(walletClient.Conn)

	r := repository.NewCardRepository(db)
	r.Initialize()
	u := usecase.NewCardUsecase(r)
	s := service.NewCardService(u, userClientService, walletClientService)

	server, _ := server.NewGrpcServer(cfg.ServerPort, s)
	server.Run()
}
