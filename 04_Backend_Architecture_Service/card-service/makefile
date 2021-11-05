all:
	go build -o card-service . ;\
	sh run.sh ;\

clean:
	rm card-service

run:
	SERVER_PORT=9083 USER_CLIENT_URL=:9081 WALLET_CLIENT_URL=:9082 POSTGRES_HOST=localhost POSTGRES_PORT=5483 POSTGRES_DB=card POSTGRES_USER=admin POSTGRES_PASSWORD=password POSTGRES_SSL_MODE=false go run main.go