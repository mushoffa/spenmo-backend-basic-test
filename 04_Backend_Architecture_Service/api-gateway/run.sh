#!/usr/bin/env bash
export SERVER_PORT=9080
export USER_CLIENT_URL=:9081 
export WALLET_CLIENT_URL=:9082 
export CARD_CLIENT_URL=:9083 
export TRANSACTION_CLIENT_URL=:9084

nohup ./api-gateway  >> api-gateway.log 2>&1 &