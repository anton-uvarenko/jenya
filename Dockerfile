FROM golang:1.22

WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]
