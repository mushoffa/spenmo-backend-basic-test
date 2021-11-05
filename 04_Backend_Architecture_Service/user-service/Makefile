all:
	go build -o user-service . ;\
	sh run.sh ;\

clean:
	rm user-service

run:
	SERVER_PORT=9081 POSTGRES_HOST=localhost POSTGRES_PORT=5481 POSTGRES_DB=user POSTGRES_USER=admin POSTGRES_PASSWORD=password POSTGRES_SSL_MODE=false go run main.go