FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /app/main .

EXPOSE 8080

CMD ["./main"]