package config

import (
	_ "github.com/kelseyhightower/envconfig"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type Config struct {
	ServerPort       int    `envconfig:"SERVER_PORT"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresSSLMode  bool   `envconfig:"POSTGRES_SSL_MODE"`
}
