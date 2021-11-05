all:
	go build -o wallet-service . ;\
	sh run.sh ;\

clean:
	rm wallet-service

run:
	SERVER_PORT=9082 USER_CLIENT_URL=:9081 POSTGRES_HOST=localhost POSTGRES_PORT=5482 POSTGRES_DB=wallet POSTGRES_USER=admin POSTGRES_PASSWORD=password POSTGRES_SSL_MODE=false go run main.go