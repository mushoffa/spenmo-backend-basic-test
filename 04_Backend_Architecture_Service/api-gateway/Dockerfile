FROM golang:1.15 AS builder
#RUN apk --no-cache 
 
WORKDIR /build
COPY ./api-gateway .
RUN go mod tidy
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o api-gateway .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/api-gateway .

EXPOSE 9080

CMD ["./api-gateway"]