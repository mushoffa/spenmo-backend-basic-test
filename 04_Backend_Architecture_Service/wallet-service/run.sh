#!/usr/bin/env bash
export SERVER_PORT=9082
export USER_CLIENT_URL=:9081
export POSTGRES_HOST=localhost 
export POSTGRES_PORT=5482 
export POSTGRES_DB=wallet 
export POSTGRES_USER=admin 
export POSTGRES_PASSWORD=password 
export POSTGRES_SSL_MODE=false

nohup ./wallet-service  >> wallet-service.log 2>&1 &