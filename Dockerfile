FROM golang:latest

WORKDIR /go/integrador_pos

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["tail", "-f", "/dev/null"]