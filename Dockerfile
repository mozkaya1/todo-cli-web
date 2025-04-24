FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o runserver server/server.go

EXPOSE 3000

CMD ["./runserver"]
