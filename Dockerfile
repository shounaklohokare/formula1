FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.* ./

RUN go mod tidy

COPY . .

RUN go build -o app

EXPOSE 8080

CMD ["./app"]
