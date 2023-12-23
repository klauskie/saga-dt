FROM golang:1.21.5

WORKDIR /dist
COPY . .

RUN go build -o saga-orchestrator ./orchestrator/cmd/main.go

ENTRYPOINT ./saga-orchestrator