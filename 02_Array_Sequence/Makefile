run:
	go run main.go

test:
	go test -coverpkg=./... -coverprofile=coverage.out ./... ;\
	go tool cover -func=coverage.out ;\
	go tool cover -html=coverage.out ;\