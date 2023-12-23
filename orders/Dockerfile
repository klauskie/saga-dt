FROM golang:1.21.5

WORKDIR /dist
COPY . .
EXPOSE 8089

RUN go build -o api-orders ./orders/cmd/main.go

ENTRYPOINT ./api-orders