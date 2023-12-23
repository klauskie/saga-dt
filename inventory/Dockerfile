FROM golang:1.21.5

WORKDIR /dist
COPY . .
EXPOSE 4040

RUN go build -o saga-inventory-ms ./inventory/cmd/main.go

ENTRYPOINT ./saga-inventory-ms