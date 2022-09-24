# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Darma Boroev"

# COPY . /projects/staff-bot
# WORKDIR /projects/staff-bot

# RUN go mod download
# FROM alpine:latest

# WORKDIR /root/

# COPY --from=0 /projects/staff-bot/bin/bot .
# COPY --from=0 /projects/staff-bot/configs configs/

# EXPOSE 80

# CMD ["./bot"]

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build -o .bin/bot cmd/bot/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# CMD ["./bot"]
