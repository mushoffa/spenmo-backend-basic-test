all:
	go build -o api-gateway . ;\
	sh run.sh ;\

clean:
	rm api-gateway

run:
	SERVER_PORT=9080 USER_CLIENT_URL=:9081 WALLET_CLIENT_URL=:9082 CARD_CLIENT_URL=:9083 TRANSACTION_CLIENT_URL=:9084  go run main.go