# Start with the official Golang image as the base image
FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go build -o app main.go

EXPOSE 8080

CMD ["/app/app"]
