#!/usr/bin/env bash
export SERVER_PORT=9084 
export USER_CLIENT_URL=:9081 
export WALLET_CLIENT_URL=:9082 
export CARD_CLIENT_URL=:9083 
export POSTGRES_HOST=localhost 
export POSTGRES_PORT=5484 
export POSTGRES_DB=transaction 
export POSTGRES_USER=admin 
export POSTGRES_PASSWORD=password 
export POSTGRES_SSL_MODE=false

nohup ./transaction-service  >> transaction-service.log 2>&1 &