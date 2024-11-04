FROM golang:1.22.7

WORKDIR /app
COPY src/ /app

RUN go build -o main main.go

EXPOSE 8080

ENTRYPOINT ["/app/main"]
