FROM golang:1.15 AS builder
#RUN apk --no-cache 

WORKDIR /build
COPY ./transaction-service .
RUN go mod tidy
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o transaction-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/transaction-service .

EXPOSE 9084

CMD ["./transaction-service"]