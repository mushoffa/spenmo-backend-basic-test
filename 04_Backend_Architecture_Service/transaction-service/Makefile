all:
	go build -o transaction-service . ;\
	sh run.sh ;\

clean:
	rm transaction-service

run:
	SERVER_PORT=9084 USER_CLIENT_URL=:9081 WALLET_CLIENT_URL=:9082 CARD_CLIENT_URL=:9083 POSTGRES_HOST=localhost POSTGRES_PORT=5484 POSTGRES_DB=transaction POSTGRES_USER=admin POSTGRES_PASSWORD=password POSTGRES_SSL_MODE=false go run main.go