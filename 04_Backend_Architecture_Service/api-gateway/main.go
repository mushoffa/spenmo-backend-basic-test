package main

import (
	"log"

	"api-gateway/config"
	v1 "api-gateway/controller/v1"
	"api-gateway/router"

	"github.com/kelseyhightower/envconfig"
	client "github.com/mushoffa/go-library/server/grpc"
	server "github.com/mushoffa/go-library/server/http"
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

	userService := protos.NewUserServiceClient(userClient.Conn)
	walletService := protos.NewWalletServiceClient(walletClient.Conn)
	cardService := protos.NewCardServiceClient(cardClient.Conn)

	v1Controller := v1.NewV1Controller(userService, walletService, cardService)

	router := router.NewRouter(v1Controller)
	router.InitializeRouter()

	server := server.NewHttpServer(cfg.ServerPort, router.Gin)
	server.Run()

}
