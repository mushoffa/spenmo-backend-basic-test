FROM golang:1.15 AS builder
#RUN apk --no-cache 

WORKDIR /build
COPY ./wallet-service .
RUN go mod tidy
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o wallet-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/wallet-service .

EXPOSE 9082

CMD ["./wallet-service"]