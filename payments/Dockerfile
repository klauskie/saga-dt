FROM golang:1.21.5

WORKDIR /dist
COPY . .
EXPOSE 4040

RUN go build -o saga-payments-ms ./payments/cmd/main.go

ENTRYPOINT ./saga-payments-ms