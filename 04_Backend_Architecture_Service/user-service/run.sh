#!/usr/bin/env bash
export SERVER_PORT=9081 
export POSTGRES_HOST=localhost 
export POSTGRES_PORT=5481 
export POSTGRES_DB=user 
export POSTGRES_USER=admin 
export POSTGRES_PASSWORD=password 
export POSTGRES_SSL_MODE=false

nohup ./user-service  >> user-service.log 2>&1 &