FROM golang:1.15 AS builder 
#RUN apk --no-cache
#FROM golang:alpine AS builder

WORKDIR /build
COPY ./user-service .
#RUN ls
RUN go mod tidy
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o user-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/user-service .

EXPOSE 9081

CMD ["./user-service"]