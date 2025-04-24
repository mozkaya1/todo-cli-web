FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o server server/server.go

EXPOSE 3000

CMD ["./server"]
