package config

import (
	_ "github.com/kelseyhightower/envconfig"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type Config struct {
	ServerPort           int    `envconfig:"SERVER_PORT"`
	UserClientURL        string `envconfig:"USER_CLIENT_URL"`
	WalletClientURL      string `envconfig:"WALLET_CLIENT_URL"`
	CardClientURL        string `envconfig:"CARD_CLIENT_URL"`
	TransactionClientURL string `envconfig:"TRANSACTION_CLIENT_URL"`
}
