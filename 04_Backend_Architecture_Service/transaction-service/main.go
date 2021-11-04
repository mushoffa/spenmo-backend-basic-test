package main

import (
	"log"
	"transaction-service/config"
	"transaction-service/data/datasource/postgres"
	"transaction-service/data/repository"
	"transaction-service/domain/usecase"
	"transaction-service/server"
	"transaction-service/service"

	"github.com/kelseyhightower/envconfig"
	client "github.com/mushoffa/go-library/server/grpc"
	"github.com/mushoffa/spenmo-proto/protos"
	"google.golang.org/grpc"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func main() {
	var cfg config.Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error loading environment config: %v", err)
	}

	db, err := postgres.NewTransactionDB(&cfg)
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

	cardClient, err := client.NewGrpcClient(cfg.CardClientURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Card GRPC Client: %s", err)
	}

	userClientService := protos.NewUserServiceClient(userClient.Conn)
	walletClientService := protos.NewWalletServiceClient(walletClient.Conn)
	cardClientService := protos.NewCardServiceClient(cardClient.Conn)

	r := repository.NewTransactionRepository(db)
	r.Initialize()
	u := usecase.NewTransactionUsecase(r)
	s := service.NewTransactionService(u, userClientService, walletClientService, cardClientService)

	server, _ := server.NewGrpcServer(cfg.ServerPort, s)
	server.Run()
}
