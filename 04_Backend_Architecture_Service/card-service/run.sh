#!/usr/bin/env bash
export SERVER_PORT=9083 
export USER_CLIENT_URL=:9081 
export WALLET_CLIENT_URL=:9082 
export POSTGRES_HOST=localhost 
export POSTGRES_PORT=5483 
export POSTGRES_DB=card 
export POSTGRES_USER=admin 
export POSTGRES_PASSWORD=password 
export POSTGRES_SSL_MODE=false

nohup ./card-service  >> card-service.log 2>&1 &