# Start with the official Golang image as the base image
FROM golang:1.23

WORKDIR /app

COPY . /app

RUN go build -o app main.go

EXPOSE 8080

# Sleep for 2 minutes (120 seconds) before starting the app
CMD sleep 80 && /app/app
